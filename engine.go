package ginger

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Engine struct {
	Gin  *gin.Engine
	Cron *cron.Cron
}

func NewEngine() *Engine {
	return &Engine{
		Gin:  gin.Default(),
		Cron: cron.New(),
	}
}

func (e *Engine) CORS(config cors.Config) {
	e.Gin.Use(cors.New(config))
}

func (e *Engine) Run(addr ...string) error {
	go e.Cron.Start()
	return e.Gin.Run(addr...)
}
