package arma

import (
	"fmt"

	pkgarma "delegacia.com.br/app/domain/arma"
)

type InsertUseCase struct {
	Service   pkgarma.Service
	Assembler *ArmaAssembler
}

type InsertUseCaseParams struct {
	Service pkgarma.Service
}

func NewInsertArmaUseCase(params InsertUseCaseParams) *InsertUseCase {
	return &InsertUseCase{
		Service: params.Service,
	}
}

func (uc *InsertUseCase) Execute() (*ArmaPresenter, error) {
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
