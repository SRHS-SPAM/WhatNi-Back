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
	r.POST("/make-student", func(c *gin.Context) {
		services.MakeUser(c, db)
	})
	r.PUT("/refresh-wh", func(c *gin.Context) {
		services.RefreshWh(c, db)
	})
	r.PUT("/wh/:hakbun", func(c *gin.Context) {
		services.Wh(c, db)
	})

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
