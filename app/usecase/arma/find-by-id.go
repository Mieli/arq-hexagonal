package arma

import (
	"fmt"

	pkgarma "delegacia.com.br/app/domain/arma"
)

type FindByIdUseCase struct {
	Service pkgarma.Service
	ID      *int64
}

type FindByIdUseCaseParams struct {
	Service pkgarma.Service
}

func NewFindByIdUseCase(params FindByIdUseCaseParams) FindByIdUseCase {
	return FindByIdUseCase{
		Service: params.Service,
	}
}

func (uc *FindByIdUseCase) Execute() (*ArmaPresenter, error) {

	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	arma, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}
	return &ArmaPresenter{
		ID:          arma.ID,
		Description: arma.Description,
		Type: TypeArmaPresenter{
			Category:    arma.Type.Category,
			Description: arma.Type.Description,
		},
	}, nil
}
