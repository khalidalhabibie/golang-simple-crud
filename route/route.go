package route

import (
	"net/http"

	"github.com/khalidalhabibie/depatu/handler"
	"github.com/khalidalhabibie/depatu/middlerware"

	"github.com/gin-gonic/gin"
)

func Run(address string) error {

	userHandler := handler.NewUserHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome Depatu")
	})

	apiRoutes := r.Group("/api/v1")
	apiRoutes.POST("/register", userHandler.AddUser)
	apiRoutes.POST("/signin", userHandler.SignInUser)

	userProtectedRoutes := apiRoutes.Group("/user", middleware.AuthorizeMiddleware())
	{
		userProtectedRoutes.GET("/profil", userHandler.GetProfile)
		userProtectedRoutes.PUT("/update", userHandler.UpdateUser)
	}

	return r.Run(address)

}
