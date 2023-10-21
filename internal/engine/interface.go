package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type IEngine interface {
	Gin() *gin.Engine
	Cron() *cron.Cron
	SystemUUID() string
	SystemName() string
}
