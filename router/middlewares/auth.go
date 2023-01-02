package middlewares

import (
	"goblog/auth"
	"goblog/error"
	"goblog/rep"

	"github.com/gin-gonic/gin"
)

func GetUserByToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(200, rep.FatalResponseWithCode(error.ErrTokenMissing))
			ctx.Abort()
			return
		}

		user, e := auth.GetUser(token)

		if e != nil {
			ctx.JSON(200, rep.BuildFatalResponse(e))
			return
		}

		ctx.Set("user", user)
		ctx.Next()

	}
}

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var usr, _ = ctx.Get("user")

		if usr != nil {
			ctx.Next()
			return
		}

		ctx.JSON(200, rep.FatalResponseWithCode(error.ErrTokenInvalid))
		ctx.Abort()
	}

}
