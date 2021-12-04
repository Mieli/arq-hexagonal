package victim_usecase

import (
	"fmt"

	pkgvictim "delegacia.com.br/app/domain/victim"
)

type UpdateUseCase struct {
	VictimService pkgvictim.Service
	Assembler     *VictimAssembler
}

type UpdateUseCaseParams struct {
	VictimService pkgvictim.Service
}

func NewUpdateUseCase(params UpdateUseCaseParams) UpdateUseCase {
	return UpdateUseCase{
		VictimService: params.VictimService,
	}
}

func (uc *UpdateUseCase) Execute() (*VictimPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("data invalid")
	}

	victim := GenerateVictim(uc.Assembler)
	newVictim, err := uc.VictimService.Save(*victim)
	if err != nil {
		return nil, fmt.Errorf("update error")
	}

	presenter := GenerateVictimPresenter(newVictim)
	return presenter, nil

}
