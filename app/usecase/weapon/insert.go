package weapon_usecase

import (
	"fmt"

	pkgweapon "delegacia.com.br/app/domain/weapon"
)

type InsertUseCase struct {
	Service   pkgweapon.Service
	Assembler *WeaponAssembler
}

type InsertUseCaseParams struct {
	Service pkgweapon.Service
}

func NewInsertUseCase(params InsertUseCaseParams) *InsertUseCase {
	return &InsertUseCase{
		Service: params.Service,
	}
}

func (uc *InsertUseCase) Execute() (*WeaponPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("Invalid data")
	}

	Weapona := GenerateWeapon(*uc.Assembler)

	newWeapona, err := uc.Service.Save(&Weapona)
	if err != nil {
		return nil, err
	}

	presenter := GenerateWeaponPresenter(*newWeapona)

	return &presenter, nil
}
