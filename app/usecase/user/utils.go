package user_usecase

import (
	pkguser "delegacia.com.br/app/domain/user"
)

func GenerateUser(assembler *UserAssembler) pkguser.User {
	user := pkguser.User{}
	if assembler != nil {
		user.ID = assembler.ID
		user.Name = assembler.Name
		user.Email = assembler.Email
	}
	return user
}

func GenerateUserPresenter(user *pkguser.User) UserPresenter {
	presenter := UserPresenter{}
	if user != nil {
		presenter.ID = user.ID
		presenter.Name = user.Name
		presenter.Email = user.Email
	}
	return presenter
}
