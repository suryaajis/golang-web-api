package model

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Profession, error)
	FindByID(ID int) (Profession, error)
	Create(profession Profession) (Profession, error)
	Update(profession Profession) (Profession, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Profession, error) {
	var professions []Profession
	err := r.db.Find(&professions).Error
	return professions, err
}

func (r *repository) FindByID(id int) (Profession, error) {
	var profession Profession
	err := r.db.Find(&profession, id).Error
	return profession, err
}

func (r *repository) Create(profession Profession) (Profession, error) {
	err := r.db.Create(&profession).Error
	return profession, err
}

func (r *repository) Update(profession Profession) (Profession, error) {
	err := r.db.Save(&profession).Error
	return profession, err
}
