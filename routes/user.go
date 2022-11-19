package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rbozburun/StudentTeachAPI/controller"
)

func UserRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/user", controller.ListUsers)
	v1.POST("/user", controller.CreateUser)

}
