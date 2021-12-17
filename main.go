package main

import (
	"golang-web-api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// API Versioning (membuat routing versi yang lain)
	v1 := router.Group("/")
	v2 := router.Group("/v2")

	v1.GET("/", handler.RootHandler)
	v1.GET("/profile", handler.ProfileHandler)
	v1.GET("/profession", handler.ProfessionHandler)
	v1.GET("/profession/:id", handler.ProfessionDetailHandler)
	v1.POST("/profession", handler.AddProfession)

	v2.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"judul": "server versi 2",
		})
	})

	// router.Run(":5000")
	router.Run()
}
