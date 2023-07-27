package ginger

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Engine struct {
	Gin        *gin.Engine
	Cron       *cron.Cron
	SystemUUID string
	SystemName string
}

func NewEngine(uuid, name string) *Engine {
	if uuid == "" {
		panic("engine uuid is empty")
	}
	if name == "" {
		panic("engine name is empty")
	}
	return &Engine{
		Gin:        gin.Default(),
		Cron:       cron.New(),
		SystemUUID: uuid,
		SystemName: name,
	}
}

func (e *Engine) Run(addr ...string) error {
	go e.Cron.Start()
	if len(addr) == 0 {
		addr = []string{"127.0.0.1:5000"}
	}
	return e.Gin.Run(addr...)
}
