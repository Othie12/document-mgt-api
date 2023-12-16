package controllers

import (
	"fmt"
	"hrms/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AddDocument(c *gin.Context) {
	userId := c.Param("user_id")
	name := c.Param("name")

	dst := "./public/"

	// Check if the "uploads" directory exists, create it if not
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = os.MkdirAll(dst, 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create 'uploads' directory"})
			return
		}
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get the file: " + err.Error()})
		return
	}

	var user models.User
	user.ID = userId
	result := models.DB.Find(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
	}

	file.Filename = fmt.Sprintf("%s_%s_%s", user.TicketNo, strings.TrimSpace(name[0:3]), time.Now())

	// Check for the existence of the file before saving
	filePath := filepath.Join(dst, file.Filename+".pdf")

	// Save the file
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save the file: " + err.Error()})
		log.Println(err.Error())
		return
	}

	document := models.Doc{
		UserID:   userId,
		Name:     name,
		Filepath: fmt.Sprintf("/%s", filePath),
	}

	models.DB.Create(&document)
	c.JSON(http.StatusOK, document)
}
