package controllers

import (
	"hrms/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexDocuments(c *gin.Context) {
	var documents []models.Document
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		panic("String convesion error" + err.Error())
	}
	models.DB.Limit(limit).Offset(offset).Find(&documents)
	c.JSON(200, documents)
}

func StoreDocuments(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	application, _ := c.FormFile("applciation")
	certificate, _ := c.FormFile("certificate")
	appointment, _ := c.FormFile("appointment")
	discipline, _ := c.FormFile("discipline")
	others, _ := c.FormFile("others")

	dst := os.Getenv("UPLOAD_PATH")

	err := c.SaveUploadedFile(application, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save application: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(certificate, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save certificate: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(appointment, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save appointment: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(discipline, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save discipline: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(others, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save others: " + err.Error()})
		return
	}

	document := models.Document{
		UserID:      id,
		Application: filepath.Join(dst, application.Filename),
		Certificate: filepath.Join(dst, certificate.Filename),
		Appointment: filepath.Join(dst, appointment.Filename),
		Discipline:  filepath.Join(dst, discipline.Filename),
		Others:      filepath.Join(dst, others.Filename),
	}

	docid := models.DB.Create(&document)
	if docid == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save data" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, document)
}

func UpdateDocument(c *gin.Context) {
	var document models.Document

	//first convert the request's id param from string to uint64
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	document.ID = uint(id) //convert id into uint from uint64 and assign it to the doc

	//get the filename (application/certificate/othier)
	name := c.Param("name")

	dst := os.Getenv("UPLOAD_PATH")
	file, _ := c.FormFile("file")

	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save to db: " + err.Error()})
		return
	}

	models.DB.Model(&document).Update(name, filepath.Join(dst, file.Filename))

	c.JSON(http.StatusOK, document)
}

func AddDocument(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	name := c.Param("name")
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to get user id: " + err.Error()})
		return
	}

	dst := os.Getenv("UPLOAD_PATH")
	file, _ := c.FormFile("file")

	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save to db: " + err.Error()})
		return
	}

	document := models.Doc{
		UserID:   uint(userId),
		Name:     name,
		Filepath: filepath.Join(dst, file.Filename),
	}

	models.DB.Create(&document)
	c.JSON(http.StatusOK, document)
}
