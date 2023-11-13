package ginger

import (
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/sql"
)

type Context[T any] struct {
	Engine    IEngine
	GinCtx    *gin.Context
	Page      *sql.Pagination
	Sort      *sql.Sort
	Request   *T
	Response  *Response
	StartTime time.Time
	IsFile    bool
	IsResp    bool
}

func (ctx *Context[T]) GetEngine() IEngine {
	return ctx.Engine
}

func (ctx *Context[T]) SetEngine(engine IEngine) {
	ctx.Engine = engine
}

func (ctx *Context[T]) GetGinCtx() *gin.Context {
	return ctx.GinCtx
}

func (ctx *Context[T]) SetGinCtx(ginCtx *gin.Context) {
	ctx.GinCtx = ginCtx
}

func (ctx *Context[T]) GetPage() *sql.Pagination {
	return ctx.Page
}

func (ctx *Context[T]) SetPage(page *sql.Pagination) {
	ctx.Page = page
}

func (ctx *Context[T]) GetSort() *sql.Sort {
	return ctx.Sort
}

func (ctx *Context[T]) SetSort(sort *sql.Sort) {
	ctx.Sort = sort
}

func (ctx *Context[T]) GetRequest() *T {
	return ctx.Request
}

func (ctx *Context[T]) SetRequest(request *T) {
	ctx.Request = request
}

func (ctx *Context[T]) GetResponse() *Response {
	return ctx.Response
}

func (ctx *Context[T]) SetResponse(response *Response) {
	ctx.Response = response
}

func (ctx *Context[T]) GetStartTime() time.Time {
	return ctx.StartTime
}

func (ctx *Context[T]) SetStartTime(startTime time.Time) {
	ctx.StartTime = startTime
}

func (ctx *Context[T]) GetIsFile() bool {
	return ctx.IsFile
}

func (ctx *Context[T]) SetIsFile(isFile bool) {
	ctx.IsFile = isFile
}

func (ctx *Context[T]) GetIsResponded() bool {
	return ctx.IsResp
}

func (ctx *Context[T]) SetIsResponded(isResp bool) {
	ctx.IsResp = isResp
}

func (ctx *Context[T]) ClientIP() string {
	return ctx.GinCtx.ClientIP()
}

func (ctx *Context[T]) UserAgent() string {
	return ctx.GinCtx.Request.UserAgent()
}

func (ctx *Context[T]) BearerToken() string {
	token := ctx.GinCtx.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	token = strings.ReplaceAll(token, "bearer ", "")
	token = strings.ReplaceAll(token, "BEARER ", "")
	token = strings.ReplaceAll(token, " ", "")
	return token
}

func (ctx *Context[T]) OK(data any, page ...*sql.Pagination) {
	if ctx.IsResp {
		log.Println("Warning: context already responded")
		return
	}
	// pagination
	var pageResponse *sql.Pagination
	if len(page) > 0 {
		pageResponse = page[0]
	}
	ctx.SetResponse(&Response{
		Success:    true,
		Pagination: pageResponse,
		Data:       data,
		Duration:   time.Since(ctx.StartTime).Milliseconds(),
	})
	ctx.SetIsResponded(true)
}

func (ctx *Context[T]) Err(code string, message string) {
	if ctx.IsResp {
		log.Println("Warning: context already responded")
		return
	}
	ctx.SetResponse(&Response{
		Success:  false,
		Duration: time.Since(ctx.StartTime).Milliseconds(),
		Error: &Error{
			Code:    code,
			Message: message,
		},
	})
	ctx.SetIsResponded(true)
}
