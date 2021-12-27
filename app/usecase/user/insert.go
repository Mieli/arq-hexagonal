package user_usecase

import (
	"fmt"

	pkguser "delegacia.com.br/app/domain/user"
	pkgservice "delegacia.com.br/infra/services"
)

type InsertUseCase struct {
	Assembler *UserAssembler
	Service   pkguser.Service
}
type InsertUseCaseParams struct {
	Service pkguser.Service
}

func NewInsertUserUseCase(params InsertUseCaseParams) InsertUseCase {
	return InsertUseCase{
		Service: params.Service,
	}
}

func (uc *InsertUseCase) Execute() (*UserPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("Invalid data")
	}
	user := pkguser.User{
		ID:    uc.Assembler.ID,
		Name:  uc.Assembler.Name,
		Email: uc.Assembler.Email,
	}
	user.Password = pkgservice.Sha256Encoder(uc.Assembler.Password)

	newUser, err := uc.Service.Save(user)
	if err != nil {
		return nil, err
	}

	presenter := GenerateUserPresenter(newUser)

	return &presenter, nil
}
