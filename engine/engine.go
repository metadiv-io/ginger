package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Engine struct {
	gin        *gin.Engine
	cron       *cron.Cron
	systemUUID string
	systemName string
}

func (e *Engine) Gin() *gin.Engine {
	return e.gin
}

func (e *Engine) Cron() *cron.Cron {
	return e.cron
}

func (e *Engine) SystemUUID() string {
	return e.systemUUID
}

func (e *Engine) SystemName() string {
	return e.systemName
}
