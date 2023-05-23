package model

import (
	"github.com/metadiv-io/ginger/err_map"
)

type Service[T any] func(ctx IContext[T]) (any, err_map.Error)
