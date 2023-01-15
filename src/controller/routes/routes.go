package routes

import (
	"github.com/EduardoPPCaldas/GoSimpleApplication/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.GET("/getUserById/:userId", controller.FindUserById)
	rg.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	rg.POST("/", controller.CreateUser)
	rg.PUT("/:userId", controller.UpdateUser)
	rg.DELETE("/:userId", controller.DeleteUser)
}