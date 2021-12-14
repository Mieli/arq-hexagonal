package victim

type VictimRepository interface {
	Save(victim Victim) (*Victim, error)
	FindById(id string) (*Victim, error)
	FindAll() ([]*Victim, error)
	Remove(id string) error
}
