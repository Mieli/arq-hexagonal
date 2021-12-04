package victim_usecase

import (
	"fmt"

	pkgvictim "delegacia.com.br/app/domain/victim"
)

type FindByIdUseCase struct {
	VictimService pkgvictim.Service
	ID            *int64
}
type FindByIdUseCaseParams struct {
	VictimService pkgvictim.Service
}

func NewFindByIdUseCase(params FindByIdUseCaseParams) *FindByIdUseCase {
	return &FindByIdUseCase{
		VictimService: params.VictimService,
	}
}

func (uc *FindByIdUseCase) Execute() (*VictimPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	victim, err := uc.VictimService.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}
	if victim == nil {
		return nil, fmt.Errorf("not found")
	}
	return &VictimPresenter{
		ID:        victim.ID,
		Name:      victim.Name,
		CPF:       victim.CPF,
		Telephone: victim.Telephone,
		Email:     victim.Email,
	}, nil

}
