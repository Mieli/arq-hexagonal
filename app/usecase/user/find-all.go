package user_usecase

import (
	pkguser "delegacia.com.br/app/domain/user"
)

type FindAllUseCase struct {
	Service pkguser.Service
}
type FindAllUseCaseParams struct {
	Service pkguser.Service
}

func NewFindAllUseCase(params FindAllUseCaseParams) FindAllUseCase {
	return FindAllUseCase{
		Service: params.Service,
	}
}

func (uc *FindAllUseCase) Execute() (*[]UserPresenter, error) {

	users, err := uc.Service.FindAll()
	if err != nil {
		return nil, err
	}
	presenters := make([]UserPresenter, 0)
	if len(users) > 0 {
		for _, user := range users {
			presenter := GenerateUserPresenter(user)

			presenters = append(presenters, presenter)

		}
	}

	return &presenters, nil

}
