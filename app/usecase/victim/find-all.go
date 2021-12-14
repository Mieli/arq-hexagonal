package victim_usecase

import (
	pkgvictim "delegacia.com.br/app/domain/victim"
)

type FindAllUseCase struct {
	VictimService pkgvictim.Service
}

type FindAllUseCaseParams struct {
	VictimService pkgvictim.Service
}

func NewFindAllUseCase(params FindAllUseCaseParams) *FindAllUseCase {
	return &FindAllUseCase{
		VictimService: params.VictimService,
	}
}

func (uc *FindAllUseCase) Execute() ([]*VictimPresenter, error) {

	result, err := uc.VictimService.FindAll()
	if err != nil {
		return nil, err
	}

	presenter := GenerateListPresenter(result)

	return presenter, nil

}
