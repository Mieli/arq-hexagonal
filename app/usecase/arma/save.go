package arma

import (
	"fmt"

	pkgarma "delegacia.com.br/app/domain/arma"
)

type SaveUseCase struct {
	Service   pkgarma.Service
	Assembler *ArmaAssembler
}

type SaveUseCaseParams struct {
	Service pkgarma.Service
}

func NewSaveArmaUseCase(params SaveUseCaseParams) *SaveUseCase {
	return &SaveUseCase{
		Service: params.Service,
	}
}

func (uc *SaveUseCase) Eecute() (*ArmaPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("Invalid data")
	}

	arma := generateArma(*uc.Assembler)

	newArma, err := uc.Service.Save(&arma)
	if err != nil {
		return nil, err
	}

	presenter := generateArmaPresenter(*newArma)

	return &presenter, nil
}

func generateArma(assembler ArmaAssembler) pkgarma.Arma {
	return pkgarma.Arma{
		ID:          assembler.ID,
		Description: assembler.Description,
		Type: pkgarma.TypeArma{
			Category:    assembler.Type.Category,
			Description: assembler.Type.Description,
		},
	}
}
func generateArmaPresenter(arma pkgarma.Arma) ArmaPresenter {
	return ArmaPresenter{
		ID:          arma.ID,
		Description: arma.Description,
		Type: TypeArmaPresenter{
			Category:    arma.Type.Category,
			Description: arma.Type.Description,
		},
	}
}
