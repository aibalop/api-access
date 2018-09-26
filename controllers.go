package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "Guest")
	password := c.PostForm("password")

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"usuario": username,
		"contraseña": password,
		"token": "123sdasd123asdasd13123sdadad"})
}
