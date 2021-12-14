package victim

type Service interface {
	Save(victim Victim) (*Victim, error)
	FindById(id string) (*Victim, error)
	FindAll() ([]*Victim, error)
	Remove(id string) error
}

type victimService struct {
	repository VictimRepository
}

func (s *victimService) Save(victim Victim) (*Victim, error) {
	return s.repository.Save(victim)
}
func (s *victimService) FindById(id string) (*Victim, error) {
	return s.repository.FindById(id)
}
func (s *victimService) FindAll() ([]*Victim, error) {
	return s.repository.FindAll()
}
func (s *victimService) Remove(id string) error {
	return s.repository.Remove(id)
}

func NewServiceVictim(repository VictimRepository) Service {
	return &victimService{
		repository: repository,
	}
}
