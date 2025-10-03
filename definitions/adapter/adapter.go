package adapter

import (
	"github.com/Nemutagk/godb/v2/definitions/config"
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
