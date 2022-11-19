package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbozburun/StudentTeachAPI/config"
	"github.com/rbozburun/StudentTeachAPI/models"
)

func ListUsers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(http.StatusOK, &users)

}

func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)

	if err := config.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusCreated, &user)
	}
}
