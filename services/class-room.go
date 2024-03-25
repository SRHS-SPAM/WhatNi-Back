package services

import (
	"WhatNi-Back/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func MakeClass(c *gin.Context, db *gorm.DB) {

}

func MakeUser(c *gin.Context, db *gorm.DB) {
	var Rstudent *entities.DTOStudent
	err := c.ShouldBindJSON(&Rstudent)
	if err != nil {
		panic(err)
	}
	student := entities.Student{
		Name:   Rstudent.Name,
		Hakbun: Rstudent.Hakbun,
		School: Rstudent.School,
		Class:  Rstudent.Class,
		Wh:     false,
	}

	err = db.Create(&student).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
