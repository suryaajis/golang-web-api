package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/profile", profileHandler)

	router.Run(":5000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Surya Aji Santosa",
		"role": "Software Engineer",
	})
}

func profileHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":   "Hello viewer",
		"content": "practice web golang web api",
	})
}
