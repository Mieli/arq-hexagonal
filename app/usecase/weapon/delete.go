package weapon_usecase

import (
	"fmt"

	pkgweapon "delegacia.com.br/app/domain/weapon"
)

type DeleteUseCase struct {
	Service pkgweapon.Service
	ID      *int64
}

type DeleteUseCaseParams struct {
	Service pkgweapon.Service
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
