package model

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/sql"
)

type IContext[T any] interface {
	GinContext() *gin.Context
	SetGinContext(*gin.Context)
	Page() *sql.Pagination
	SetPage(*sql.Pagination)
	Sort() *sql.Sort
	SetSort(*sql.Sort)
	Request() *T
	SetRequest(*T)
	Response() any
	SetResponse(any)
	ClientIP() string
	UserAgent() string
}

type Context[T any] struct {
	ginContext *gin.Context
	page       *sql.Pagination
	sort       *sql.Sort
	request    *T
	response   any
}

func (ctx *Context[T]) GinContext() *gin.Context {
	return ctx.ginContext
}

func (ctx *Context[T]) SetGinContext(ginContext *gin.Context) {
	ctx.ginContext = ginContext
}

func (ctx *Context[T]) Page() *sql.Pagination {
	return ctx.page
}

func (ctx *Context[T]) SetPage(page *sql.Pagination) {
	ctx.page = page
}

func (ctx *Context[T]) Sort() *sql.Sort {
	return ctx.sort
}

func (ctx *Context[T]) SetSort(sort *sql.Sort) {
	ctx.sort = sort
}

func (ctx *Context[T]) Request() *T {
	return ctx.request
}

func (ctx *Context[T]) SetRequest(request *T) {
	ctx.request = request
}

func (ctx *Context[T]) Response() any {
	return ctx.response
}

func (ctx *Context[T]) SetResponse(response any) {
	ctx.response = response
}

func (ctx *Context[T]) ClientIP() string {
	return ctx.ginContext.ClientIP()
}

func (ctx *Context[T]) UserAgent() string {
	return ctx.ginContext.Request.UserAgent()
}
