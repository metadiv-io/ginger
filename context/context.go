package context

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger/engine"
	"github.com/metadiv-io/ginger/internal/constant"
	"github.com/metadiv-io/ginger/internal/err_map"
	"github.com/metadiv-io/ginger/internal/util"
	"github.com/metadiv-io/ginger/types"
	"github.com/metadiv-io/sql"
)

type Context[T any] struct {
	engine engine.IEngine
	ginCtx *gin.Context
	page   *sql.Pagination
	sort   *sql.Sort
	req    *T
	resp   *types.Response
	start  time.Time
	isFile bool
	isResp bool
}

func (ctx *Context[T]) Engine() engine.IEngine {
	return ctx.engine
}

func (ctx *Context[T]) GinCtx() *gin.Context {
	return ctx.ginCtx
}

func (ctx *Context[T]) Page() *sql.Pagination {
	return ctx.page
}

func (ctx *Context[T]) Sort() *sql.Sort {
	return ctx.sort
}

func (ctx *Context[T]) Request() *T {
	return ctx.req
}

func (ctx *Context[T]) Response() *types.Response {
	return ctx.resp
}

func (ctx *Context[T]) SetResponse(resp *types.Response) {
	ctx.resp = resp
}

func (ctx *Context[T]) StartTime() time.Time {
	return ctx.start
}

func (ctx *Context[T]) IsFile() bool {
	return ctx.isFile
}

func (ctx *Context[T]) IsResponded() bool {
	return ctx.isResp
}

func (ctx *Context[T]) SetIsResponded(isResp bool) {
	ctx.isResp = isResp
}

func (ctx *Context[T]) ClientIP() string {
	return ctx.GinCtx().ClientIP()
}

func (ctx *Context[T]) UserAgent() string {
	return ctx.GinCtx().Request.UserAgent()
}

func (ctx *Context[T]) SetTraceID(traceID string) {
	ctx.GinCtx().Header(constant.HEADER_TRACE_ID, traceID)
}

func (ctx *Context[T]) TraceID() string {
	traceID := ctx.GinCtx().GetHeader(constant.HEADER_TRACE_ID)
	if traceID == "" {
		traceID = util.NewNanoID("Trace-")
		ctx.SetTraceID(traceID)
	}
	return traceID
}

func (ctx *Context[T]) Locale() string {
	return ctx.GinCtx().GetHeader(constant.HEADER_LOCALE)
}

func (ctx *Context[T]) SetLocale(locale string) {
	ctx.GinCtx().Header(constant.HEADER_LOCALE, locale)
}

func (ctx *Context[T]) Traces() []types.Trace {
	traces := make([]types.Trace, 0)
	tracesHeader := ctx.GinCtx().GetHeader(constant.HEADER_TRACE)
	if tracesHeader != "" {
		_ = json.Unmarshal([]byte(tracesHeader), &traces)
	}
	return traces
}

func (ctx *Context[T]) SetTraces(traces []types.Trace) {
	tracesBytes, _ := json.Marshal(traces)
	ctx.GinCtx().Header(constant.HEADER_TRACE, string(tracesBytes))
}

func (ctx *Context[T]) BearerToken() string {
	token := ctx.GinCtx().GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	token = strings.ReplaceAll(token, "bearer ", "")
	token = strings.ReplaceAll(token, "BEARER ", "")
	token = strings.ReplaceAll(token, " ", "")
	return token
}

func (ctx *Context[T]) OK(data any, page ...*sql.Pagination) {
	if ctx.IsResponded() {
		log.Println("Warning: context already responded")
		return
	}

	// handle traces
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Success:    true,
		SystemUUID: ctx.Engine().SystemUUID(),
		SystemName: ctx.Engine().SystemName(),
		TraceID:    ctx.TraceID(),
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		Duration:   time.Since(ctx.StartTime()).Milliseconds(),
	})

	// pagination
	var pageResponse *sql.Pagination
	if len(page) > 0 {
		pageResponse = page[0]
	}

	ctx.SetResponse(&types.Response{
		Success:    true,
		TraceID:    ctx.TraceID(),
		Locale:     ctx.Locale(),
		Duration:   time.Since(ctx.StartTime()).Milliseconds(),
		Pagination: pageResponse,
		Data:       data,
		Traces:     traces,
	})
	ctx.Response().Calculate()
	ctx.SetIsResponded(true)
}

func (ctx *Context[T]) Err(code string) {
	if ctx.IsResponded() {
		log.Println("Warning: context already responded")
		return
	}

	// handle traces
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Success:    false,
		SystemUUID: ctx.Engine().SystemUUID(),
		SystemName: ctx.Engine().SystemName(),
		TraceID:    ctx.TraceID(),
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		Duration:   time.Since(ctx.StartTime()).Milliseconds(),
		Error:      types.NewError(code, err_map.ErrMap.Get(code, ctx.Locale())),
	})

	ctx.SetResponse(&types.Response{
		Success:  false,
		TraceID:  ctx.TraceID(),
		Locale:   ctx.Locale(),
		Duration: time.Since(ctx.StartTime()).Milliseconds(),
		Traces:   traces,
		Error:    types.NewError(code, err_map.ErrMap.Get(code, ctx.Locale())),
	})
	ctx.Response().Calculate()
	ctx.SetIsResponded(true)
}
