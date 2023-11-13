package ginger

import "github.com/metadiv-io/sql"

type Response struct {
	Success    bool            `json:"success"`
	Duration   int64           `json:"duration"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Error      *Error          `json:"error,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
}
