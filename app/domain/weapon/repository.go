package weapon

type WeaponRepository interface {
	Save(Weapona *Weapon) (*Weapon, error)
	FindById(id int64) (*Weapon, error)
	FindAll() (*[]Weapon, error)
	Remove(id int64) error
}
