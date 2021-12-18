package db

import (
	"golang-web-api/pkg/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/profession-api"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Failed")
	}

	// MIGRATE TABLE
	db.AutoMigrate(&model.Profession{})

	// CREATE
	// profession := model.Profession{}
	// profession.Name = "Police"
	// profession.Salary = 3500000
	// profession.Rating = 9
	// profession.Description = "People who secure country"
	// err = db.Create(&profession).Error
	// if err != nil {
	// 	fmt.Println("Error create profession record")
	// }

	// READ
	/*
		1. db.First (data id pertama)
		2. db.Last (data id terakhir)
		3. db.First(model, primarykey) ambil berdasarkan primary key
		4. db.Find (untuk mengambil semua data) bentuk slice / array
		# https://gorm.io/docs/query.html
	*/
	// var profession model.Profession
	// err = db.Debug().Where("id = ?", 1).First(&profession).Error
	// if err != nil {
	// 	fmt.Println("Error read profession record")
	// }
	// fmt.Println(profession)

	// UPDATE
	// get / read data first and then update it
	// var profession model.Profession
	// err = db.Debug().Where("id = ?", 1).First(&profession).Error
	// if err != nil {
	// 	fmt.Println("Error read profession record")
	// }
	// profession.Rating = 9
	// err = db.Save(&profession).Error
	// if err != nil {
	// 	fmt.Println("Error update profession record")
	// }

	// DELETE
	// var profession model.Profession
	// err = db.Debug().Where("id = ?", 1).First(&profession).Error
	// if err != nil {
	// 	fmt.Println("Error read profession record")
	// }
	// err = db.Delete(&profession).Error
	// if err != nil {
	// 	fmt.Println("Error update profession record")
	// }

	// IMPORT CRUD Function from Model (Repository)
	// professionRepository := model.NewRepository(db)
	// MICROSERVICES CRUD FUNCTION
	// professionService := model.NewService(professionRepository)

	// IMPLEMENT
	// professions, err := professionRepository.FindAll()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, data := range professions {
	// 		fmt.Println(data)
	// 	}
	// }

	// profession, err := professionRepository.FindByID(2)
	// fmt.Println(profession)

	// input := model.Profession{
	// 	Name:        "Sailor",
	// 	Salary:      20000000,
	// 	Rating:      8,
	// 	Description: "People who sail ship",
	// }
	// profession, err := professionRepository.Create(input)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(profession)

	// professionService.Create(model.ProfessionRequest{
	// 	Name:        "Farmer",
	// 	Salary:      "2000000",
	// 	Rating:      7,
	// 	Description: "People who take care rice field",
	// })

	return db
}
