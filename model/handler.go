package model

import (
	"github.com/gin-gonic/gin"
)

type Handler[T any] func() Service[T]

type HandlerOptions struct {
	BeforeService *gin.HandlerFunc
	AfterService  *gin.HandlerFunc
}
