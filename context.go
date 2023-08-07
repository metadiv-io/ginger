package ginger

import (
	"encoding/json"
	"log"
	"strings"
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
	IsFile      bool
	IsResponded bool
}

func NewContext[T any](engine *Engine, ginCtx *gin.Context) *Context[T] {
	var page *sql.Pagination
	var sort *sql.Sort
	if ginCtx.Request.Method == "GET" {
		page = GinRequest[sql.Pagination](ginCtx)
		sort = GinRequest[sql.Sort](ginCtx)
	}
	ctx := &Context[T]{
		Engine:      engine,
		GinCtx:      ginCtx,
		Page:        page,
		Sort:        sort,
		Request:     GinRequest[T](ginCtx),
		Response:    nil,
		StartTime:   time.Now(),
		IsResponded: false,
	}
	ctx.TraceID() // generate trace id if not exist
	return ctx
}

func (ctx *Context[T]) ClientIP() string {
	return ctx.GinCtx.ClientIP()
}

func (ctx *Context[T]) UserAgent() string {
	return ctx.GinCtx.Request.UserAgent()
}

func (ctx *Context[T]) TraceID() string {
	traceID := ctx.GinCtx.GetHeader(HEADER_TRACE_ID)
	if traceID == "" {
		traceID = NewNanoID("Trace-")
		ctx.SetTraceID(traceID)
	}
	return traceID
}

func (ctx *Context[T]) SetTraceID(traceID string) {
	ctx.GinCtx.Header(HEADER_TRACE_ID, traceID)
}

func (ctx *Context[T]) Locale() string {
	return ctx.GinCtx.GetHeader(HEADER_LOCALE)
}

func (ctx *Context[T]) SetLocale(locale string) {
	ctx.GinCtx.Header(HEADER_LOCALE, locale)
}

func (ctx *Context[T]) Traces() []types.Trace {
	traces := make([]types.Trace, 0)
	tracesHeader := ctx.GinCtx.GetHeader(HEADER_TRACE)
	if tracesHeader != "" {
		_ = json.Unmarshal([]byte(tracesHeader), &traces)
	}
	return traces
}

func (ctx *Context[T]) SetTraces(traces []types.Trace) {
	tracesBytes, _ := json.Marshal(traces)
	ctx.GinCtx.Header(HEADER_TRACE, string(tracesBytes))
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
	if ctx.IsResponded {
		log.Println("Warning: context already responded")
		return
	}

	// handle traces
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Success:    true,
		SystemUUID: ctx.Engine.SystemUUID,
		SystemName: ctx.Engine.SystemName,
		TraceID:    ctx.TraceID(),
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		Duration:   time.Since(ctx.StartTime).Milliseconds(),
	})

	// pagination
	var pageResponse *sql.Pagination
	if len(page) > 0 {
		pageResponse = page[0]
	}

	ctx.Response = &types.Response{
		Success:    true,
		TraceID:    ctx.TraceID(),
		Locale:     ctx.Locale(),
		Duration:   time.Since(ctx.StartTime).Milliseconds(),
		Pagination: pageResponse,
		Data:       data,
		Traces:     traces,
	}
	ctx.Response.Calculate()
	ctx.IsResponded = true
}

func (ctx *Context[T]) Err(code string) {
	if ctx.IsResponded {
		log.Println("Warning: context already responded")
		return
	}

	// handle traces
	traces := ctx.Traces()
	traces = append(traces, types.Trace{
		Success:    false,
		SystemUUID: ctx.Engine.SystemUUID,
		SystemName: ctx.Engine.SystemName,
		TraceID:    ctx.TraceID(),
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		Duration:   time.Since(ctx.StartTime).Milliseconds(),
		Error:      types.NewError(code, ErrMap.Get(code, ctx.Locale())),
	})

	ctx.Response = &types.Response{
		Success:  false,
		TraceID:  ctx.TraceID(),
		Locale:   ctx.Locale(),
		Duration: time.Since(ctx.StartTime).Milliseconds(),
		Traces:   traces,
		Error:    types.NewError(code, ErrMap.Get(code, ctx.Locale())),
	}
	ctx.Response.Calculate()
	ctx.IsResponded = true
}

func (ctx *Context[T]) OKDownloadFile(bytes []byte, filename ...string) {
	if ctx.IsResponded {
		log.Println("Warning: context already responded")
		return
	}
	name := ctx.determineFilename(filename...)
	ctx.GinCtx.Header("Content-Disposition", "attachment; filename="+name)
	ctx.GinCtx.Data(200, "application/octet-stream", bytes)
	ctx.IsFile = true
	ctx.IsResponded = true
}

func (ctx *Context[T]) OKServeFile(bytes []byte, filename ...string) {
	if ctx.IsResponded {
		log.Println("Warning: context already responded")
		return
	}
	name := ctx.determineFilename(filename...)
	ctx.GinCtx.Header("Content-Disposition", "attachment; filename="+name)
	ctx.GinCtx.Data(200, ctx.determineContentType(name), bytes)
	ctx.IsFile = true
	ctx.IsResponded = true
}

func (ctx *Context[T]) determineFilename(filename ...string) string {
	if len(filename) == 0 {
		return NewNanoID()
	}
	if filename[0] == "" {
		return NewNanoID()
	}
	return filename[0]
}

func (ctx *Context[T]) determineContentType(filename string) string {
	var ext string
	xs := strings.Split(filename, ".")
	if len(xs) > 1 {
		ext = xs[len(xs)-1]
	}

	switch ext {
	case "png", "jpg", "jpeg", "gif", "bmp", "webp", "svg", "ico":
		return "image/" + ext
	case "pdf":
		return "application/pdf"
	case "doc", "docx":
		return "application/msword"
	case "xls", "xlsx":
		return "application/vnd.ms-excel"
	case "ppt", "pptx":
		return "application/vnd.ms-powerpoint"
	case "zip":
		return "application/zip"
	case "mp3":
		return "audio/mpeg"
	case "mp4":
		return "video/mp4"
	case "txt":
		return "text/plain"
	default:
		return "application/octet-stream"
	}
}
