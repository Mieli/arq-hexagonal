package user_usecase

import (
	"fmt"

	pkguser "delegacia.com.br/app/domain/user"
)

type FindByIDUseCase struct {
	Service pkguser.Service
	ID      string
}
type FindByIDUseCaseParams struct {
	Service pkguser.Service
}

func NewFindByIdUseCase(params FindByIDUseCaseParams) FindByIDUseCase {
	return FindByIDUseCase{
		Service: params.Service,
	}
}

func (uc *FindByIDUseCase) Execute() (*UserPresenter, error) {
	if uc.ID == "" {
		return nil, fmt.Errorf("Invalid data")
	}
	user, err := uc.Service.FindById(uc.ID)
	if err != nil {
		return nil, err
	}

	presenter := GenerateUserPresenter(user)

	return &presenter, nil

}
