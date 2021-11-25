package arma

import (
	pkgarma "delegacia.com.br/app/domain/arma"
)

type FindAllUseCase struct {
	Service pkgarma.Service
}
type FindAllUseCaseParams struct {
	Service pkgarma.Service
}

func NewFindAllUseCase(params FindAllUseCaseParams) FindAllUseCase {
	return FindAllUseCase{
		Service: params.Service,
	}
}

func (uc *FindAllUseCase) Execute() (*[]ArmaPresenter, error) {
	armas, err := uc.Service.FindAll()
	if err != nil {
		return nil, err
	}
	presenter := generatePresenter(armas)

	return &presenter, nil
}

func generatePresenter(armas *[]pkgarma.Arma) []ArmaPresenter {
	list := make([]ArmaPresenter, 0)
	if armas != nil && len(*armas) > 0 {
		for _, arma := range *armas {
			list = append(list, ArmaPresenter{
				ID:          arma.ID,
				Description: arma.Description,
				Type: TypeArmaPresenter{
					Category:    arma.Type.Category,
					Description: arma.Type.Description,
				},
			})
		}
	}
	return list
}
