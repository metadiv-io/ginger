package ginger

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func CRON(engine IEngine, spec string, job func()) {
	engine.GetCron().AddFunc(spec, job)
}

func GET[T any](engine IEngine, route string, handler ApiHandler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().GET(route, append(middleware, handler.GinHandler(engine))...)
}

func GETCached[T any](engine IEngine, route string, handler ApiHandler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GetGin().GET(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler.GinHandler(engine)))...)
}

func POST[T any](engine IEngine, route string, handler ApiHandler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().POST(route, append(middleware, handler.GinHandler(engine))...)
}

func POSTCached[T any](engine IEngine, route string, handler ApiHandler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GetGin().POST(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler.GinHandler(engine)))...)
}

func PUT[T any](engine IEngine, route string, handler ApiHandler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().PUT(route, append(middleware, handler.GinHandler(engine))...)
}

func PUTCached[T any](engine IEngine, route string, handler ApiHandler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GetGin().PUT(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler.GinHandler(engine)))...)
}

func DELETE[T any](engine IEngine, route string, handler ApiHandler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().DELETE(route, append(middleware, handler.GinHandler(engine))...)
}

func DELETECached[T any](engine IEngine, route string, handler ApiHandler[T], duration time.Duration, middleware ...gin.HandlerFunc) {
	engine.GetGin().DELETE(route, append(middleware, cache.CachePage(persistence.NewInMemoryStore(time.Second), duration, handler.GinHandler(engine)))...)
}

func WS[T any](engine IEngine, route string, handler WsHandler[T], middleware ...gin.HandlerFunc) {
	engine.GetGin().GET(route, append(middleware, handler.GinHandler(engine))...)
}
