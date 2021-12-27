package user_usecase

import (
	"fmt"

	pkguser "delegacia.com.br/app/domain/user"
)

type DeleteUseCase struct {
	Service pkguser.Service
	ID      *string
}

type DeleteUseCaseParams struct {
	Service pkguser.Service
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
