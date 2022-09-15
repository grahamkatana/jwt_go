package routes

import {
	"os"
	"github.com/gin-gonic/gin"
	controller "golang_jwt/controllers"

}

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup",controller.Signup())
	incomingRoutes.POST("users/login",controller.Login())

}