package ginger

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger/engine"
	"github.com/robfig/cron"
)

type IEngine engine.IEngine

type Engine struct {
	Gin        *gin.Engine
	Cron       *cron.Cron
	SystemUUID string
	SystemName string
}

var NewEngine = engine.NewEngine
