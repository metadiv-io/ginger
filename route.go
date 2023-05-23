package ginger

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginger/model"
	"github.com/metadiv-io/sql"
)

func GET[T any](engine *Engine, route string, handler model.Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.GET(route, append(middleware, handlerFunc(handler))...)
}

func GETCached[T any](engine *Engine, route string, handler model.Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GinEngine.GET(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handlerFunc(handler)))...)
}

func POST[T any](engine *Engine, route string, handler model.Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.POST(route, append(middleware, handlerFunc(handler))...)
}

func PUT[T any](engine *Engine, route string, handler model.Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.PUT(route, append(middleware, handlerFunc(handler))...)
}

func DELETE[T any](engine *Engine, route string, handler model.Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.DELETE(route, append(middleware, handlerFunc(handler))...)
}

func handlerFunc[T any](h model.Handler[T]) gin.HandlerFunc {
	service := h()
	return func(ctx *gin.Context) {

		now := time.Now()

		c := new(model.Context[T])
		c.SetGinContext(ctx)
		c.SetRequest(ParseRequest[T](ctx))
		c.SetPage(ParseRequest[sql.Pagination](ctx))
		c.SetSort(ParseRequest[sql.Sort](ctx))

		resp, err := service(c)

		if err != nil {
			respObj := &model.Response{
				Success:  false,
				Time:     time.Now().Format(time.RFC3339),
				Duration: time.Since(now).Milliseconds(),
				Error: model.Error{
					Code:    err.Code(),
					Message: err.Error(),
				},
			}
			c.SetResponse(respObj)
			ctx.JSON(http.StatusOK, respObj)
			return
		}

		respObj := &model.Response{
			Success:    true,
			Time:       time.Now().Format(time.RFC3339),
			Duration:   time.Since(now).Milliseconds(),
			Data:       resp,
			Pagination: c.Page(),
		}
		c.SetResponse(respObj)
		ctx.JSON(http.StatusOK, respObj)
	}
}
