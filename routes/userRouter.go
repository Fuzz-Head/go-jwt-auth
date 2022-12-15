package routes

import (
	controller "github.com/fuzzhead/controllers"
	middleware "github.com/fuzzhead/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.Getusers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}
