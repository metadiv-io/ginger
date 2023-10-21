package ginger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger/internal/context"
	"github.com/metadiv-io/ginger/internal/engine"
	"github.com/metadiv-io/ginger/internal/util"
	"github.com/metadiv-io/sql"
)

type IContext[T any] context.IContext[T]

func NewContext[T any](e engine.IEngine, ginCtx *gin.Context) IContext[T] {
	var page *sql.Pagination
	var sort *sql.Sort
	if ginCtx.Request.Method == "GET" {
		page = new(sql.Pagination)
		sort = new(sql.Sort)
		ginCtx.ShouldBindQuery(page)
		ginCtx.ShouldBindQuery(sort)
	}
	ctx := context.NewContext[T](
		e,
		ginCtx,
		page,
		sort,
		util.GinRequest[T](ginCtx),
		nil,
		time.Now(),
		false,
		false,
	)
	ctx.TraceID() // generate trace id if not exist
	return ctx
}
