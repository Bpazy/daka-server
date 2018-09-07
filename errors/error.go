package errors

import (
	"github.com/Bpazy/daka-server/util"
	"github.com/gin-gonic/gin"
)

func New(s string) error {
	return &Error{s}
}

type Error struct {
	S string
}

func (e Error) Error() string {
	return e.S
}

func JsonError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		last := c.Errors.Last()
		if last == nil {
			return
		}
		e, ok := last.Err.(*Error)
		if ok {
			c.AbortWithStatusJSON(520, util.Result{
				Code: util.FAILED,
				Msg:  e.S,
				Data: nil,
			})
		}
	}
}
