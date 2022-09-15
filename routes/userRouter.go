package routes

import {
	"os"
	"github.com/gin-gonic/gin"
	"golang_jwt/middleware"
	controller "golang_jwt/controllers"

}

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())

}