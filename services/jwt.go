package services

import (
	"WhatNi-Back/entities"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

var jwtKey = []byte("asdifhiuascbhchddjbhc") // JWT 서명을 위한 키

// Claims 구조체는 jwt.StandardClaims을 확장합니다.
type Claims struct {
	Hakbun string `json:"hakbun"`
	jwt.StandardClaims
}

// GenerateJWT 함수는 사용자 학번을 받아 JWT 토큰을 생성하고 반환합니다.
func GenerateJWT(hakbun string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 토큰 유효 시간 설정
	claims := &Claims{
		Hakbun: hakbun,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func MakeUser(c *gin.Context, db *gorm.DB) {
	var Rstudent *entities.DTOStudent
	err := c.ShouldBindJSON(&Rstudent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	hakbunStr := strconv.Itoa(student.Hakbun) // int 타입의 Hakbun을 string으로 변환합니다.

	// 사용자 등록 성공 후 JWT 토큰 생성
	token, err := GenerateJWT(hakbunStr) // 변환된 문자열을 사용하여 토큰을 생성합니다.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}
