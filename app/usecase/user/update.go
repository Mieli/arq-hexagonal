package user_usecase

import (
	"fmt"

	pkguser "delegacia.com.br/app/domain/user"
	pkgservice "delegacia.com.br/infra/services"
)

type UpdateUseCase struct {
	Service   pkguser.Service
	Assembler *UserAssembler
}

type UpdateUseCaseParams struct {
	Service pkguser.Service
}

func NewUpdateUseCase(params UpdateUseCaseParams) UpdateUseCase {
	return UpdateUseCase{
		Service: params.Service,
	}
}

func (uc *UpdateUseCase) Execute() (*UserPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("data invalid")
	}

	user := GenerateUser(uc.Assembler)
	if uc.Assembler.Password != "" {
		user.Password = pkgservice.Sha256Encoder(uc.Assembler.Password)
	}
	newUser, err := uc.Service.Save(user)
	if err != nil {
		return nil, fmt.Errorf("update error")
	}

	presenter := GenerateUserPresenter(newUser)
	return &presenter, nil

}
