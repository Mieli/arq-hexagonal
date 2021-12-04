package weapon_usecase

import (
	pkgweapon "delegacia.com.br/app/domain/weapon"
)

func GenerateWeapon(assembler WeaponAssembler) pkgweapon.Weapon {
	return pkgweapon.Weapon{
		ID:          assembler.ID,
		Description: assembler.Description,
		Type: pkgweapon.TypeWeapon{
			Category:    assembler.Type.Category,
			Description: assembler.Type.Description,
		},
	}
}
func GenerateWeaponPresenter(Weapona pkgweapon.Weapon) WeaponPresenter {
	return WeaponPresenter{
		ID:          Weapona.ID,
		Description: Weapona.Description,
		Type: TypeWeaponPresenter{
			Category:    Weapona.Type.Category,
			Description: Weapona.Type.Description,
		},
	}
}

func GeneratePresenter(weapons *[]pkgweapon.Weapon) []WeaponPresenter {
	list := make([]WeaponPresenter, 0)
	if weapons != nil && len(*weapons) > 0 {
		for _, arma := range *weapons {
			list = append(list, WeaponPresenter{
				ID:          arma.ID,
				Description: arma.Description,
				Type: TypeWeaponPresenter{
					Category:    arma.Type.Category,
					Description: arma.Type.Description,
				},
			})
		}
	}
	return list
}
