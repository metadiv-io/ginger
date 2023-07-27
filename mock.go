package ginger

import (
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/sql"
)

type MockContextParams[T any] struct {
	Request   *T
	Page      *sql.Pagination
	Sort      *sql.Sort
	Method    string
	Path      string
	ClientIP  string
	UserAgent string
	Headers   map[string]string
}

func MockContext[T any](params MockContextParams[T]) *Context[T] {
	w := httptest.NewRecorder()
	ctx, e := gin.CreateTestContext(w)

	if params.Method == "" {
		params.Method = "GET"
	}
	if params.Path == "" {
		params.Path = "/"
	}

	ctx.Request = httptest.NewRequest(params.Method, params.Path, nil)

	if params.ClientIP != "" {
		ctx.Request.RemoteAddr = params.ClientIP
	} else {
		ctx.Request.RemoteAddr = "127.0.0.1"
	}

	if params.UserAgent != "" {
		ctx.Request.Header.Set("User-Agent", params.UserAgent)
	} else {
		ctx.Request.Header.Set("User-Agent", "Mock Context")
	}

	if params.Headers != nil {
		for k, v := range params.Headers {
			ctx.Request.Header.Set(k, v)
		}
	}

	return &Context[T]{
		GinCtx: ctx,
		Engine: &Engine{
			Gin: e,
		},
		Request:   params.Request,
		Page:      params.Page,
		Sort:      params.Sort,
		StartTime: time.Now(),
	}
}
