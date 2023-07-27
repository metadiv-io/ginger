package ginger

import "github.com/gin-gonic/gin"

type Service[T any] func(ctx *Context[T])

type Handler[T any] func() Service[T]

func (h Handler[T]) GinHandler(engine *Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler := h()
		c := NewContext[T](engine, ctx)
		handler(c)

		// unexpected, service did not respond
		if !c.IsResponded || c.Response == nil {
			ctx.JSON(500, gin.H{
				"message": "service did not respond",
			})
			return
		}

		// success
		if c.Response.Success {
			ctx.JSON(200, c.Response)
			return
		}

		// error, but no error object
		if c.Response.Error == nil {
			ctx.JSON(500, gin.H{
				"message": "service did not respond with error",
			})
			return
		}
	}
}
