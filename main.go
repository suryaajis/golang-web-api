package main

import (
	"golang-web-api/handler"
	"golang-web-api/pkg/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=profession-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db.Init()

	if err != nil {
		log.Fatal("DB Connection Failed")
	}
	db.AutoMigrate(&model.Profession{})

	profesiRepository := model.NewRepository(db)
	profesiService := model.NewService(profesiRepository)
	profesiHandler := handler.NewProfessionHandler(profesiService)

	router := gin.Default()

	// API Versioning (membuat routing versi yang lain)
	v1 := router.Group("/")
	v2 := router.Group("/v2")

	v1.GET("/", profesiHandler.RootHandler)
	v1.GET("/profile", profesiHandler.ProfileHandler)
	v1.GET("/profession", profesiHandler.GetProfessions)
	v1.GET("/profession/:id", profesiHandler.GetProfessionDetail)
	v1.POST("/profession", profesiHandler.AddProfession)
	v1.PUT("/profession/:id", profesiHandler.EditProfession)

	v2.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"judul": "server versi 2",
		})
	})

	// router.Run(":5000")
	router.Run()
}
