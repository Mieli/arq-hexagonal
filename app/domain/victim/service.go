package victim

type Service interface {
	Save(victim Victim) (*Victim, error)
	FindById(id int64) (*Victim, error)
	FindAll() (*[]Victim, error)
	Remove(id int64) error
}

type victimSerive struct {
	repository VictimRepository
}

func (s *victimSerive) Save(victim Victim) (*Victim, error) {
	return s.repository.Save(victim)
}
func (s *victimSerive) FindById(id int64) (*Victim, error) {
	return s.repository.FindById(id)
}
func (s *victimSerive) FindAll() (*[]Victim, error) {
	return s.repository.FindAll()
}
func (s *victimSerive) Remove(id int64) error {
	return s.repository.Remove(id)
}

func NewServiceVictim(repository VictimRepository) Service {
	return &victimSerive{
		repository: repository,
	}
}
