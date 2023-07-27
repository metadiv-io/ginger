package ginger

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func GET[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}

func GETCached[T any](engine *Engine, route string, handler Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.Gin.GET(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler.GinHandler(engine)))...)
}

func POST[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.Gin.POST(route, append(middleware, handler.GinHandler(engine))...)
}

func POSTCached[T any](engine *Engine, route string, handler Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.Gin.POST(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler.GinHandler(engine)))...)
}

func PUT[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.Gin.PUT(route, append(middleware, handler.GinHandler(engine))...)
}

func PUTCached[T any](engine *Engine, route string, handler Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.Gin.PUT(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler.GinHandler(engine)))...)
}

func DELETE[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.Gin.DELETE(route, append(middleware, handler.GinHandler(engine))...)
}

func DELETECached[T any](engine *Engine, route string, handler Handler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.Gin.DELETE(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler.GinHandler(engine)))...)
}

func WS[T any](engine *Engine, route string, handler WsHandler[T], middleware ...gin.HandlerFunc) {
	engine.Gin.GET(route, append(middleware, handler.GinHandler(engine))...)
}
