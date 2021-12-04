package victim

type VictimRepository interface {
	Save(victim Victim) (*Victim, error)
	FindById(id int64) (*Victim, error)
	FindAll() (*[]Victim, error)
	Remove(id int64) error
}
