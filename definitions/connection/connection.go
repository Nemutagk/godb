package connection

import (
	"github.com/Nemutagk/godb/v2/definitions/adapter"
	"github.com/Nemutagk/godb/v2/definitions/config"
)

type Connection struct {
	Adapter adapter.Adapter
	Config  config.Config
}
