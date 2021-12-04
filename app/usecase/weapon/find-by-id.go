package weapon_usecase

import (
	"fmt"

	pkgweapon "delegacia.com.br/app/domain/weapon"
)

type FindByIdUseCase struct {
	Service pkgweapon.Service
	ID      *int64
}

type FindByIdUseCaseParams struct {
	Service pkgweapon.Service
}

func NewFindByIdUseCase(params FindByIdUseCaseParams) FindByIdUseCase {
	return FindByIdUseCase{
		Service: params.Service,
	}
}

func (uc *FindByIdUseCase) Execute() (*WeaponPresenter, error) {

	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	weapon, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}
	return &WeaponPresenter{
		ID:          weapon.ID,
		Description: weapon.Description,
		Type: TypeWeaponPresenter{
			Category:    weapon.Type.Category,
			Description: weapon.Type.Description,
		},
	}, nil
}
