package caller

import (
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/sql"
)

type Response[T any] struct {
	Success    bool            `json:"success"`
	Duration   int64           `json:"duration"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Error      *ginger.Error   `json:"error,omitempty"`
	Data       *T              `json:"data,omitempty"`
}
