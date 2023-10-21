package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func NewEngine(uuid, name string) IEngine {
	return &Engine{
		gin:        gin.Default(),
		cron:       cron.New(),
		systemUUID: uuid,
		systemName: name,
	}
}

func NewMockEngine(e *gin.Engine) IEngine {
	return &Engine{
		gin:        e,
		cron:       cron.New(),
		systemUUID: "",
		systemName: "",
	}
}
