package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rbozburun/StudentTeachAPI/config"
	"github.com/rbozburun/StudentTeachAPI/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":4545")
}
