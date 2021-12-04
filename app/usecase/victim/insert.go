package victim_usecase

import (
	"fmt"

	pkgvictim "delegacia.com.br/app/domain/victim"
)

type InsertUseCase struct {
	VictimService pkgvictim.Service
	Assembler     *VictimAssembler
}

type InsertUseCaseParams struct {
	VictimService pkgvictim.Service
}

func NewInsertUseCase(params InsertUseCaseParams) *InsertUseCase {
	return &InsertUseCase{
		VictimService: params.VictimService,
	}
}

func (uc *InsertUseCase) Execute() (*VictimPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}
	victim := GenerateVictim(uc.Assembler)
	result, err := uc.VictimService.Save(*victim)
	if err != nil {
		return nil, err
	}
	presenter := GenerateVictimPresenter(result)
	return presenter, nil
}
