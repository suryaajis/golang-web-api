package model

type Service interface {
	FindAll() ([]Profession, error)
	FindByID(ID int) (Profession, error)
	Create(professionRequest ProfessionRequest) (Profession, error)
	Update(ID int, professionRequest ProfessionRequest) (Profession, error)
	Delete(ID int) (Profession, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Profession, error) {
	professions, err := s.repository.FindAll()
	return professions, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(id int) (Profession, error) {
	profession, err := s.repository.FindByID(id)
	return profession, err
	// return s.repository.FindByID(id)
}

func (s *service) Create(professionRequest ProfessionRequest) (Profession, error) {
	salary, _ := professionRequest.Salary.Int64()

	profesi := Profession{
		Name:        professionRequest.Name,
		Salary:      int(salary),
		Rating:      professionRequest.Rating,
		Description: professionRequest.Description,
	}

	newProfesi, err := s.repository.Create(profesi)
	return newProfesi, err
}

func (s *service) Update(ID int, professionRequest ProfessionRequest) (Profession, error) {
	profesi, _ := s.repository.FindByID(ID)

	salary, _ := professionRequest.Salary.Int64()

	profesi.Name = professionRequest.Name
	profesi.Salary = int(salary)
	profesi.Rating = professionRequest.Rating
	profesi.Description = professionRequest.Description

	newProfesi, err := s.repository.Update(profesi)
	return newProfesi, err
}

func (s *service) Delete(ID int) (Profession, error) {
	profesi, _ := s.repository.FindByID(ID)
	deleteProfesi, err := s.repository.Delete(profesi)
	return deleteProfesi, err
}
