package weapon_usecase

import (
	pkgweapon "delegacia.com.br/app/domain/weapon"
)

type FindAllUseCase struct {
	Service pkgweapon.Service
}
type FindAllUseCaseParams struct {
	Service pkgweapon.Service
}

func NewFindAllUseCase(params FindAllUseCaseParams) FindAllUseCase {
	return FindAllUseCase{
		Service: params.Service,
	}
}

func (uc *FindAllUseCase) Execute() (*[]WeaponPresenter, error) {
	armas, err := uc.Service.FindAll()
	if err != nil {
		return nil, err
	}
	presenter := GeneratePresenter(armas)

	return &presenter, nil
}
