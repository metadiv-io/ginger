package ginger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/sql"
)

func NewContext[T any](
	engine IEngine,
	ginCtx *gin.Context,
) IContext[T] {
	var page *sql.Pagination
	var sort *sql.Sort
	if ginCtx.Request.Method == "GET" {
		page = new(sql.Pagination)
		sort = new(sql.Sort)
		ginCtx.ShouldBindQuery(page)
		ginCtx.ShouldBindQuery(sort)
	}
	ctx := &Context[T]{
		Engine:    engine,
		GinCtx:    ginCtx,
		Page:      page,
		Sort:      sort,
		Request:   ginRequest[T](ginCtx),
		StartTime: time.Now(),
	}
	return ctx
}

type IContext[T any] interface {
	GetEngine() IEngine
	SetEngine(engine IEngine)
	GetGinCtx() *gin.Context
	SetGinCtx(ginCtx *gin.Context)

	GetPage() *sql.Pagination
	SetPage(page *sql.Pagination)
	GetSort() *sql.Sort
	SetSort(sort *sql.Sort)
	GetResponse() *Response
	SetResponse(resp *Response)
	GetRequest() *T
	SetRequest(req *T)

	GetStartTime() time.Time
	SetStartTime(startTime time.Time)
	GetIsFile() bool
	SetIsFile(isFile bool)
	GetIsResponded() bool
	SetIsResponded(isResp bool)

	ClientIP() string
	UserAgent() string
	BearerToken() string
}
