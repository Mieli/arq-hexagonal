package weapon

type Service interface {
	Save(Weapon *Weapon) (*Weapon, error)
	FindById(id int64) (*Weapon, error)
	FindAll() (*[]Weapon, error)
	Remove(id int64) error
}

type WeaponService struct {
	repository WeaponRepository
}

func (s *WeaponService) Save(Weapon *Weapon) (*Weapon, error) {
	return s.repository.Save(Weapon)
}
func (s *WeaponService) FindById(id int64) (*Weapon, error) {
	return s.repository.FindById(id)
}
func (s *WeaponService) FindAll() (*[]Weapon, error) {
	return s.repository.FindAll()
}
func (s *WeaponService) Remove(id int64) error {
	return s.repository.Remove(id)
}

func NewServiceWeapon(repository WeaponRepository) Service {
	return &WeaponService{
		repository: repository,
	}
}
