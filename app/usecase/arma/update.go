package arma

import (
	"fmt"

	pkgarma "delegacia.com.br/app/domain/arma"
)

type UpdateUseCase struct {
	Service   pkgarma.Service
	Assembler *ArmaAssembler
}

type UpdateUseCaseParams struct {
	Service pkgarma.Service
}

func NewUpdateUseCase(params UpdateUseCaseParams) UpdateUseCase {
	return UpdateUseCase{
		Service: params.Service,
	}
}

func (uc *UpdateUseCase) Execute() (*ArmaPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("data invalid")
	}

	arma := generateArma(*uc.Assembler)
	newArma, err := uc.Service.Save(&arma)
	if err != nil {
		return nil, fmt.Errorf("update error")
	}

	presenter := generateArmaPresenter(*newArma)
	return &presenter, nil

}
