package route

import (
	"net/http"

	"golang-simple-crud/handler"
	middleware "golang-simple-crud/middlerware"

	"github.com/gin-gonic/gin"
)

func Run(address string) error {

	userHandler := handler.NewUserHandler()
	taskHandler := handler.NewTaskHandler()
	addressHandler := handler.NewAddressHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome Depatu")
	})
	r.Static("/file", ".")

	apiRoutes := r.Group("/api/v1")
	apiRoutes.POST("/register", userHandler.AddUser)
	apiRoutes.POST("/signin", userHandler.SignInUser)

	userProtectedRoutes := apiRoutes.Group("/user", middleware.AuthorizeMiddleware())
	{
		userProtectedRoutes.GET("/profil", userHandler.GetProfile)
		userProtectedRoutes.PUT("/profile", userHandler.UpdateUser)
		userProtectedRoutes.POST("/address", addressHandler.AddAddress)
		userProtectedRoutes.GET("/address", addressHandler.GetAddressbyUser)
		userProtectedRoutes.PUT("/address", addressHandler.UpdateAddress)
		userProtectedRoutes.DELETE("/address", addressHandler.DeleteAddress)
		userProtectedRoutes.POST("/photo", userHandler.UploadPhoto)
		userProtectedRoutes.PUT("/password", userHandler.UpdatePassword)

	}

	adminProtectedRoutes := apiRoutes.Group("/admin", middleware.AuthorizeAdminMiddleware())
	{
		adminProtectedRoutes.GET("/users", userHandler.GetAllUser)
		adminProtectedRoutes.DELETE("/task", taskHandler.DeleteTask)
		adminProtectedRoutes.GET("/tasks", taskHandler.GetAllTask)
		adminProtectedRoutes.POST("/task/user", taskHandler.AddTask)
		adminProtectedRoutes.GET("/task/user", taskHandler.GetTaskUserAdmin)
		adminProtectedRoutes.GET("/address/user", addressHandler.GetAddressUserAdmin)

	}

	return r.Run(address)
}
