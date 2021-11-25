package arma

type ArmaRepository interface {
	Save(arma *Arma) (*Arma, error)
	FindById(id int64) (*Arma, error)
	FindAll() (*[]Arma, error)
	Remove(id int64) error
}
