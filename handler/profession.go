package handler

import (
	"fmt"
	"golang-web-api/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type professionHandler struct {
	professionService model.Service
}

func NewProfessionHandler(professionService model.Service) *professionHandler {
	return &professionHandler{professionService}
}

func (h *professionHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Surya Aji Santosa",
		"role": "Software Engineer",
	})
}

func (h *professionHandler) ProfileHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":   "Hello viewer",
		"content": "practice web golang web api",
	})
}

func (h *professionHandler) GetProfessions(c *gin.Context) {
	// get params endpoint/path
	// name := c.Query("name")
	// category := c.Query("category")

	profesi, err := h.professionService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": profesi,
	})
}

func (h *professionHandler) GetProfessionDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	response, err := h.professionService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *professionHandler) AddProfession(c *gin.Context) {
	// title, salary, description
	var input model.ProfessionRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// log.Fatal(err)
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	profesi, err := h.professionService.Create(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":   profesi.Name,
		"salary": profesi.Salary,
		"rating": profesi.Rating,
		"desc":   profesi.Description,
	})
}

func (h *professionHandler) EditProfession(c *gin.Context) {
	var input model.ProfessionRequest
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// log.Fatal(err)
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	profesi, err := h.professionService.Update(id, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":   profesi.Name,
		"salary": profesi.Salary,
		"rating": profesi.Rating,
		"desc":   profesi.Description,
	})
}

func (h *professionHandler) DeleteProfession(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	_, err := h.professionService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "profession has been deleted",
	})
}
