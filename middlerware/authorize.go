package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/khalidalhabibie/depatu/handler"
)

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})

		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := handler.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)

			} else {
				if token.Valid {
					ctx.Set("userID", claims["userID"])
					ctx.Set("userName", claims["userName"])
					ctx.Set("admin", claims["admin"])
					fmt.Println("during authorization", claims["userID"], claims["userName"], claims["admin"])
					if claims["admin"] != false {
						ctx.AbortWithStatus(http.StatusUnauthorized)
					}
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}

			}
		}
	}
}

func AuthorizeAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})

		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := handler.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)

			} else {
				if token.Valid {
					ctx.Set("userID", claims["userID"])
					ctx.Set("userName", claims["userName"])
					ctx.Set("admin", claims["admin"])
					fmt.Println("during authorization", claims["userID"], claims["userName"], claims["admin"])
					if claims["admin"] != true {
						ctx.AbortWithStatus(http.StatusUnauthorized)
					}

				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}

			}
		}
	}
}
