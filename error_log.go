package errorlog

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/golang/glog"
)

func ErrorResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}

		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors[len(c.Errors)-1]
		render.WriteJSON(
			c.Writer,
			gin.H{
				"error": err.Error(),
			},
		)
		glog.Error(err)
	}
}
