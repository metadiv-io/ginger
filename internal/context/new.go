package context

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger/internal/engine"
	"github.com/metadiv-io/ginger/types"
	"github.com/metadiv-io/sql"
)

func NewContext[T any](
	engine engine.IEngine,
	ginCtx *gin.Context,
	page *sql.Pagination,
	sort *sql.Sort,
	req *T,
	resp *types.Response,
	start time.Time,
	isFile bool,
	isResp bool,
) IContext[T] {
	return &Context[T]{
		engine: engine,
		ginCtx: ginCtx,
		page:   page,
		sort:   sort,
		req:    req,
		resp:   resp,
		start:  start,
		isFile: isFile,
		isResp: isResp,
	}
}
