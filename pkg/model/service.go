package model

type Service interface {
	FindAll() ([]Profession, error)
	FindByID(ID int) (Profession, error)
	Create(professionRequest ProfessionRequest) (Profession, error)
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
