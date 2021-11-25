package arma

type Service interface {
	Save(arma *Arma) (*Arma, error)
	FindById(id int64) (*Arma, error)
	FindAll() (*[]Arma, error)
	Remove(id int64) error
}

type armaService struct {
	repository ArmaRepository
}

func (s *armaService) Save(arma *Arma) (*Arma, error) {
	return s.repository.Save(arma)
}
func (s *armaService) FindById(id int64) (*Arma, error) {
	return s.repository.FindById(id)
}
func (s *armaService) FindAll() (*[]Arma, error) {
	return s.repository.FindAll()
}
func (s *armaService) Remove(id int64) error {
	return s.repository.Remove(id)
}

func NewServiceArma(repository ArmaRepository) Service {
	return &armaService{
		repository: repository,
	}
}
