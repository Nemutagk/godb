package godb

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/Nemutagk/godb/definitions/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var listConnectionStabilizedOnce sync.Once
var connectionManagerBuil *ConnectionManager

type ConnectionManager struct {
	Connections map[string]interface{}
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		Connections: make(map[string]interface{}),
	}
}

func (cm *ConnectionManager) AddConnection(name string, connection any) {
	cm.Connections[name] = connection
}

func (cm *ConnectionManager) GetConnection(name string) (any, error) {
	connection, exists := cm.Connections[name]
	if !exists {
		return nil, errors.New("connection not found")
	}
	return connection, nil
}

func InitConnections(connections map[string]db.DbConnection) *ConnectionManager {
	listConnectionStabilizedOnce.Do(func() {
		connectionManager := NewConnectionManager()

		for name, connection := range connections {
			switch connection.Driver {
			case "mongo", "mongodb":
				conn, err := mongoConnection(connection)
				if err != nil {
					panic(fmt.Errorf("failed to connect to MongoDB: %w", err))
				}

				connectionManager.AddConnection(name, conn)
			default:
				panic(fmt.Errorf("unsupported connection type: %s", connection.Driver))
			}
		}

		connectionManagerBuil = connectionManager
	})

	return connectionManagerBuil
}

func mongoConnection(connConfig db.DbConnection) (*mongo.Client, error) {
	// Check if the environment variables are set
	if connConfig.Host == "" || connConfig.Port == "" || connConfig.User == "" || connConfig.Password == "" || connConfig.Database == "" {
		panic("missing required environment variables for MongoDB connection")
	}

	if connConfig.AnotherConfig == nil {
		// fmt.Println("anotherConfig not found, setting default value")
		connConfig.AnotherConfig = &map[string]interface{}{
			"authSource": "admin",
		}
	}

	if _, ok := (*connConfig.AnotherConfig)["authSource"]; !ok {
		// fmt.Println("db_auth not found in anotherConfig, setting default value")
		(*connConfig.AnotherConfig)["authSource"] = "admin"
	}

	if _, ok := (*connConfig.AnotherConfig)["retryWrites"]; !ok {
		// fmt.Println("db_auth not found in anotherConfig, setting default value")
		(*connConfig.AnotherConfig)["retryWrites"] = "true"
	}

	if _, ok := (*connConfig.AnotherConfig)["w"]; !ok {
		// fmt.Println("db_auth not found in anotherConfig, setting default value")
		(*connConfig.AnotherConfig)["w"] = "majority"
	}

	var mongoUri string
	if _, ok := (*connConfig.AnotherConfig)["cluster"]; !ok {
		mongoUri = "mongodb://" + connConfig.Host + ":" + connConfig.Port + "/" + connConfig.Database // + "?authSource=" + (*connConfig.AnotherConfig)["db_auth"].(string)
	} else {
		mongoUri = "mongodb+srv://" + connConfig.User + ":" + connConfig.Password + "@" + connConfig.Host + "/" + connConfig.Database // + "?authSource=" + (*connConfig.AnotherConfig)["db_auth"].(string)
	}

	if connConfig.AnotherConfig != nil {
		mongoUri = mongoUri + "?"
		for key, value := range *connConfig.AnotherConfig {
			if key == "cluster" || key == "authSource" {
				continue
			}

			mongoUri = mongoUri + key + "=" + fmt.Sprintf("%v", value) + "&"
		}
		mongoUri = mongoUri[:len(mongoUri)-1]
	}

	cred := options.Credential{
		Username:   connConfig.User,
		Password:   connConfig.Password,
		AuthSource: (*connConfig.AnotherConfig)["authSource"].(string),
		// Forzar mecanismo si hay problemas:
		// AuthMechanism: "SCRAM-SHA-256",
	}

	clientOpts := options.Client().ApplyURI(mongoUri).SetAuth(cred)

	// fmt.Println("MongoDB URI:", mongoUri)
	connection, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	return connection, nil
}
