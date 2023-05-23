package model

import (
	"github.com/metadiv-io/sql"
)

type Response struct {
	Success    bool            `json:"success"`
	Time       string          `json:"time"`
	Duration   int64           `json:"duration"`
	Data       any             `json:"data,omitempty"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Error      Error           `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
