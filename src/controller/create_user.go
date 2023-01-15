package controller

import (
	"github.com/EduardoPPCaldas/GoSimpleApplication/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	err := rest_err.NewInternalServerError("Not implemented yet")
	ctx.JSON(err.Code, err)
}
