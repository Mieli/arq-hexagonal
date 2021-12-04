package weapon_usecase

import (
	"fmt"

	pkgweapon "delegacia.com.br/app/domain/weapon"
)

type UpdateUseCase struct {
	Service   pkgweapon.Service
	Assembler *WeaponAssembler
}

type UpdateUseCaseParams struct {
	Service pkgweapon.Service
}

func NewUpdateUseCase(params UpdateUseCaseParams) UpdateUseCase {
	return UpdateUseCase{
		Service: params.Service,
	}
}

func (uc *UpdateUseCase) Execute() (*WeaponPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("data invalid")
	}

	arma := GenerateWeapon(*uc.Assembler)
	newArma, err := uc.Service.Save(&arma)
	if err != nil {
		return nil, fmt.Errorf("update error")
	}

	presenter := GenerateWeaponPresenter(*newArma)
	return &presenter, nil

}
