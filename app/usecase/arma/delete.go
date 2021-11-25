package arma

import (
	"fmt"

	pkgarma "delegacia.com.br/app/domain/arma"
)

type DeleteUseCase struct {
	Service pkgarma.Service
	ID      *int64
}

type DeleteUseCaseParams struct {
	Service pkgarma.Service
}

func NewDeleteUseCase(params DeleteUseCaseParams) DeleteUseCase {
	return DeleteUseCase{
		Service: params.Service,
	}
}

func (uc *DeleteUseCase) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return fmt.Errorf("error remove data")
	}
	return nil
}
