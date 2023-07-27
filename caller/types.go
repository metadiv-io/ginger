package caller

import (
	"github.com/metadiv-io/ginger/types"
	"github.com/metadiv-io/sql"
)

type Response[T any] struct {
	Success    bool            `json:"success"`
	TraceID    string          `json:"trace_id"`
	Locale     string          `json:"locale"`
	Duration   int64           `json:"duration"`
	Credit     float64         `json:"credit"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Error      *types.Error    `json:"error,omitempty"`
	Data       *T              `json:"data,omitempty"`
	Traces     []types.Trace   `json:"traces,omitempty"`
}
