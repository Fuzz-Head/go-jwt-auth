package routes

import (
	controller "jwt-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/sigup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
