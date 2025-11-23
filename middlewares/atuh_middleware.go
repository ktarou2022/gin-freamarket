package middlewares

import (
	"gin-fleamarket/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authServiece services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headler := ctx.GetHeader("Authorization")
		if headler == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(headler, "Beare ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(headler, "Bear ")
		user, err := authServiece.GetUserFromToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
