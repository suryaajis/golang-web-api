package profession

import "encoding/json"

type Profession struct {
	Title       string      `json:"title" binding:"required"`
	Salary      json.Number `json:"salary" binding:"required,number"`
	Category    string
	Description string `json:"desc"`
}
