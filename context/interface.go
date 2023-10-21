package context

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger/engine"
	"github.com/metadiv-io/ginger/types"
	"github.com/metadiv-io/sql"
)

type IContext[T any] interface {
	Engine() engine.IEngine
	GinCtx() *gin.Context
	Page() *sql.Pagination
	Sort() *sql.Sort
	Request() *T
	Response() *types.Response
	SetResponse(resp *types.Response)
	StartTime() time.Time
	IsFile() bool
	IsResponded() bool
	SetIsResponded(isResp bool)
	ClientIP() string
	UserAgent() string
	SetTraceID(traceID string)
	TraceID() string
	Locale() string
	SetLocale(locale string)
	Traces() []types.Trace
	SetTraces(traces []types.Trace)
	BearerToken() string
	OK(data any, page ...*sql.Pagination)
	Err(code string)
}
