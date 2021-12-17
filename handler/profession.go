package handler

import (
	"fmt"
	"golang-web-api/profession"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Surya Aji Santosa",
		"role": "Software Engineer",
	})
}

func ProfileHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":   "Hello viewer",
		"content": "practice web golang web api",
	})
}

func ProfessionHandler(c *gin.Context) {
	// get params endpoint/path
	name := c.Query("name")
	category := c.Query("category")

	c.JSON(http.StatusOK, gin.H{
		"title":        "profession page",
		"profesi_name": name,
		"category":     category,
	})
}

func ProfessionDetailHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"profesi_id": id,
		"title":      "Profil Page",
		"content":    "Your full data",
	})
}

func AddProfession(c *gin.Context) {
	// title, salary, description
	var input profession.Profession

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
