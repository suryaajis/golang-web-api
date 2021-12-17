package model

import "encoding/json"

type InputProfession struct {
	Title       string      `json:"title" binding:"required"`
	Salary      json.Number `json:"salary" binding:"required,number"`
	Category    string
	Description string `json:"desc"`
}
