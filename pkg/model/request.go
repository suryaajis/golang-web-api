package model

import "encoding/json"

type ProfessionRequest struct {
	Name        string      `json:"name" binding:"required"`
	Salary      json.Number `json:"salary" binding:"required,number"`
	Rating      int
	Description string `json:"desc"`
}
