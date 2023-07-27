package ginger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger/types"
	"github.com/metadiv-io/sql"
)

type Context[T any] struct {
	Engine      *Engine
	GinCtx      *gin.Context
	Page        *sql.Pagination
	Sort        *sql.Sort
	Request     *T
	Response    *types.Response
	StartTime   time.Time
	IsResponded bool
}

func NewContext[T any](engine *Engine, ginCtx *gin.Context) *Context[T] {
	var page *sql.Pagination
	var sort *sql.Sort
	if ginCtx.Request.Method == "GET" {
		page = GinRequest[sql.Pagination](ginCtx)
		sort = GinRequest[sql.Sort](ginCtx)
	}
	return &Context[T]{
		Engine:      engine,
		GinCtx:      ginCtx,
		Page:        page,
		Sort:        sort,
		Request:     GinRequest[T](ginCtx),
		Response:    nil,
		StartTime:   time.Now(),
		IsResponded: false,
	}
}
