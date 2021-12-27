package user

type Service interface {
	FindAll() ([]*User, error)
	FindById(id string) (*User, error)
	FindByEmailAndPassord(email, password string) (*User, error)
	Save(user User) (*User, error)
	Remove(id string) error
}

type userService struct {
	repository UserRepository
}

func (s *userService) Save(user User) (*User, error) {
	return s.repository.Save(user)
}

func (s *userService) FindAll() ([]*User, error) {
	return s.repository.FindAll()
}

func (s *userService) FindById(id string) (*User, error) {
	return s.repository.FindById(id)
}
func (s *userService) Remove(id string) error {
	return s.repository.Remove(id)
}
func (s *userService) FindByEmailAndPassord(email, password string) (*User, error) {
	return s.repository.FindByEmailAndPassord(email, password)
}

func NewServiceUser(repository UserRepository) Service {
	return &userService{
		repository: repository,
	}
}
