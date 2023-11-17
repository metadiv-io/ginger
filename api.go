package ginger

import "github.com/gin-gonic/gin"

type ApiService[T any] func(ctx IContext[T])

type ApiHandler[T any] func() ApiService[T]

func (h ApiHandler[T]) GinHandler(engine IEngine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler := h()
		c := NewContext[T](engine, ctx)
		handler(c)

		// if file is served, no need to respond
		if c.GetIsResponded() && c.GetIsFile() {
			return
		}

		// unexpected, service did not respond
		if !c.GetIsResponded() || c.GetResponse() == nil {
			ctx.JSON(500, gin.H{
				"message": "service did not respond",
			})
			return
		}

		// success
		if c.GetResponse().Success {
			ctx.JSON(200, c.GetResponse())
			return
		}

		// error, but no error object
		if c.GetResponse().Error == nil {
			ctx.JSON(500, gin.H{
				"message": "service did not respond with error",
			})
			return
		}

		// error
		ctx.JSON(400, c.GetResponse())
	}
}
