package services

import (
	"WhatNi-Back/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func RefreshWh(c *gin.Context, db *gorm.DB) {
	hakbunParam := c.Param("hakbun")         // URL에서 hakbun 파라미터 값을 받아옵니다.
	hakbun, err := strconv.Atoi(hakbunParam) // hakbun 값을 int로 변환합니다.
	if err != nil {
		// 파라미터 변환 에러 처리
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 학번 포맷",
		})
		return
	}

	// 특정 hakbun을 가진 Student의 wh 값을 false로 업데이트합니다.
	err = db.Model(&entities.Student{}).Where("hakbun = ?", hakbun).Update("wh", false).Error
	if err != nil {
		// 에러 처리
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "데이터베이스 업데이트 중 오류 발생",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wh 값이 성공적으로 업데이트 되었습니다.",
	})
}

func Wh(c *gin.Context, db *gorm.DB) {
	hakbunParam := c.Param("hakbun")         // URL에서 hakbun 파라미터 값을 받아옵니다.
	hakbun, err := strconv.Atoi(hakbunParam) // hakbun 값을 int로 변환합니다.
	if err != nil {
		// 파라미터 변환 에러 처리
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 학번 포맷",
		})
		return
	}

	// 특정 hakbun을 가진 Student의 wh 값을 true로 업데이트합니다.
	err = db.Model(&entities.Student{}).Where("hakbun = ?", hakbun).Update("wh", true).Error
	if err != nil {
		// 에러 처리
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "데이터베이스 업데이트 중 오류 발생",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wh 값이 성공적으로 업데이트 되었습니다.",
	})
}
