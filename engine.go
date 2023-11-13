package ginger

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Engine struct {
	Gin  *gin.Engine
	Cron *cron.Cron
}

func (e *Engine) GetGin() *gin.Engine {
	return e.Gin
}

func (e *Engine) GetCron() *cron.Cron {
	return e.Cron
}

func (e *Engine) SetGin(g *gin.Engine) {
	e.Gin = g
}

func (e *Engine) SetCron(c *cron.Cron) {
	e.Cron = c
}
