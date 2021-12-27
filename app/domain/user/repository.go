package user

type UserRepository interface {
	FindAll() ([]*User, error)
	FindById(id string) (*User, error)
	FindByEmailAndPassord(email, password string) (*User, error)
	Save(user User) (*User, error)
	Remove(id string) error
}
