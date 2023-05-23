package ginger

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Engine struct {
	GinEngine  *gin.Engine
	CronEngine *cron.Cron
}

func (e *Engine) CORS(config cors.Config) {
	e.GinEngine.Use(cors.New(config))
}

func NewEngine() *Engine {
	return &Engine{
		GinEngine:  gin.Default(),
		CronEngine: cron.New(),
	}
}

func (e *Engine) Run(host string) {
	e.CronEngine.Start()
	e.GinEngine.Run(host)
}
