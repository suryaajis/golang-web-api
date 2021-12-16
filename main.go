package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/profile", profileHandler)
	router.GET("/profession", professionHandler)
	router.GET("/profession/:id", professionDetailHandler)
	router.POST("/profession", addProfession)

	// router.Run(":5000")
	router.Run()
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

func professionHandler(c *gin.Context) {
	// get params endpoint/path
	name := c.Query("name")
	category := c.Query("category")

	c.JSON(http.StatusOK, gin.H{
		"title":        "profession page",
		"profesi_name": name,
		"category":     category,
	})
}

func professionDetailHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"profesi_id": id,
		"title":      "Profil Page",
		"content":    "Your full data",
	})
}

type Profession struct {
	Title       string      `json:"title" binding:"required"`
	Salary      json.Number `json:"salary" binding:"required,number"`
	Category    string
	Description string `json:"desc"`
}

func addProfession(c *gin.Context) {
	// title, salary, description
	var input Profession

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// log.Fatal(err)
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"title":    input.Title,
		"salary":   input.Salary,
		"category": input.Category,
		"desc":     input.Description,
	})
}
