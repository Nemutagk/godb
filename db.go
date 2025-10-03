package godb

import (
	"errors"
	"log"
	"sync"

	"github.com/Nemutagk/godb/v2/definitions/connection"
)

var ErrConnectionNotFound = errors.New("connection not found")

var listConnectionStabilizedOnce sync.Once
var connectionManagerBuil *connectionManager

type connectionManager struct {
	Connections map[string]*connectionWrapper
}

type connectionWrapper struct {
	Connection connection.Connection
	Status     string
}

const ConnectionWrapperStatusActive = "active"
const ConnectionWrapperStatusInactive = "inactive"
const ConnectionWrapperStatusClose = "close"
const ConnectionWrapperStatusError = "error"

func NewConnectionManager() *connectionManager {
	listConnectionStabilizedOnce.Do(func() {
		connectionManagerBuil = &connectionManager{
			Connections: make(map[string]*connectionWrapper),
		}
	})

	return connectionManagerBuil
}

func (cm *connectionManager) GetConnection(name string) (*connectionWrapper, error) {
	log.Println("Getting connection: " + name)
	if conn, ok := cm.Connections[name]; ok {
		if conn.Status == ConnectionWrapperStatusActive {
			return conn, nil
		}

		if err := conn.Connection.Adapter.Connect(); err != nil {
			conn.Status = ConnectionWrapperStatusError
			cm.Connections[name] = conn
			return nil, err
		}

		conn.Status = ConnectionWrapperStatusActive
		cm.Connections[name] = conn

		return conn, nil
	}

	log.Println("Connection not found: " + name)
	return nil, ErrConnectionNotFound
}

func (cm *connectionManager) AddConnection(name string, conn connection.Connection) {
	log.Println("Adding connection: " + name)
	cm.Connections[name] = &connectionWrapper{Connection: conn, Status: ConnectionWrapperStatusInactive}
	log.Println("list connections: ", cm.Connections)
}

func (cm *connectionManager) RemoveConnection(name string) {
	if conn, ok := cm.Connections[name]; ok {
		conn.Connection.Adapter.Close()
		conn.Status = ConnectionWrapperStatusClose
		cm.Connections[name] = conn
	}
}

func (cm *connectionManager) ListConnections() map[string]*connectionWrapper {
	return cm.Connections
}

func (cm *connectionManager) LoadMultipleConnection(listConnections map[string]*connection.Connection) {
	log.Println("Loading multiple connections...")
	for i, conn := range listConnections {
		log.Println("Loading connection: " + i)
		cm.Connections[i] = &connectionWrapper{Connection: *conn, Status: ConnectionWrapperStatusInactive}
	}
}
