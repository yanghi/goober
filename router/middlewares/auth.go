package middlewares

import (
	"goblog/auth"
	"goblog/error"
	"goblog/rep"

	"github.com/gin-gonic/gin"
)

func GetUserByToken() gin.HandlerFunc {
	return getUserByToken(true)
}
func TryGetUserByToken() gin.HandlerFunc {
	return getUserByToken(false)
}

func getUserByToken(requiredToken bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			if requiredToken {
				ctx.JSON(200, rep.FatalResponseWithCode(error.ErrTokenMissing))
				ctx.Abort()
			}
			return
		}

		user, e := auth.GetUser(token)

		if e != nil {
			if requiredToken {
				ctx.JSON(200, rep.FatalResponseWithGError(*e))
				ctx.Abort()
			}
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
