package controllers

import (
	"hrms/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var users []models.User
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		panic("String conversion Error: " + err.Error())
	}
	models.DB.Limit(20).Offset(offset).Find(&users)
	c.JSON(200, users)
}
