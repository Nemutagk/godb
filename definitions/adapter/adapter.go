package adapter

import (
	"github.com/Nemutagk/godb/definitions/config"
)

type Config struct {
	Dsn  string
	Conn any
}

type Adapter interface {
	SetConf(conf config.Config) error
	Connect() error
	GetConnection() any
	Close() error
}
