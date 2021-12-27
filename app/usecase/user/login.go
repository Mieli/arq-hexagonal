package user_usecase

import (
	"fmt"

	pkguser "delegacia.com.br/app/domain/user"
	pkgservice "delegacia.com.br/infra/services"
)

type LoginUseCase struct {
	Service   pkguser.Service
	Assembler *LoginAssembler
}
type LoginUseCaseParams struct {
	Service pkguser.Service
}

func NewLoginUseCase(params LoginUseCaseParams) LoginUseCase {
	return LoginUseCase{
		Service: params.Service,
	}
}

func (uc *LoginUseCase) Execute() (*UserPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("Invalid data")
	}

	email := uc.Assembler.Email
	password := pkgservice.Sha256Encoder(uc.Assembler.Password)

	user, err := uc.Service.FindByEmailAndPassord(email, password)
	if err != nil {
		return nil, err
	}

	return &UserPresenter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
