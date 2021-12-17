package model

import "time"

type Profession struct {
	ID          int
	Name        string
	Salary      int
	Rating      int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
