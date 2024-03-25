package contorllers

import (
	"WhatNi-Back/repositories"
	"WhatNi-Back/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewContorllers(port string) {
	r := gin.Default()

	db := repositories.MySQLInit()

	r.POST("/make-room", func(c *gin.Context) {
		services.MakeClass(c, db)
	})
	r.POST("/make-student", func(c *gin.Context) {
		services.MakeUser(c, db)
	})
	r.PUT("/refresh-wh", func(c *gin.Context) {
		services.RefreshWh(c, db)
	})
	r.PUT("/wh/:hakbun", func(c *gin.Context) {
		services.Wh(c, db)
	})
	r.GET("/generate-token", func(c *gin.Context) {
		username, exists := c.GetQuery("username")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username query parameter is required"})
			return
		}

		token, err := services.GenerateJWT(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
