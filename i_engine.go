package ginger

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func NewEngine() IEngine {
	return &Engine{
		Gin:  gin.New(),
		Cron: cron.New(),
	}
}

type IEngine interface {
	GetGin() *gin.Engine
	GetCron() *cron.Cron
	SetGin(g *gin.Engine)
	SetCron(c *cron.Cron)
}
