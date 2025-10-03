package connection

import (
	"github.com/Nemutagk/godb/definitions/adapter"
	"github.com/Nemutagk/godb/definitions/config"
)

type Connection struct {
	Adapter adapter.Adapter
	Config  config.Config
}
