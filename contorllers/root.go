package contorllers

import (
	"WhatNi-Back/repositories"
	"WhatNi-Back/services"
	"github.com/gin-gonic/gin"
)

func NewContorllers(port string) {
	r := gin.Default()

	db := repositories.MySQLInit()

	r.POST("/make-room", func(c *gin.Context) {
		services.MakeClass(c, db)
	})

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
