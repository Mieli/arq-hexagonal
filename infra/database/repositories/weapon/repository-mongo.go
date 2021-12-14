package weapon

import (
	"fmt"

	pkgweapon "delegacia.com.br/app/domain/weapon"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	dataBase   *mongo.Client
	collection string
}

func NewWeaponRepository(db mongo.Client) pkgweapon.WeaponRepository {
	return &repository{
		dataBase:   &db,
		collection: "weapon",
	}
}

func (r *repository) Save(weapon *pkgweapon.Weapon) (*pkgweapon.Weapon, error) {
	newWeapon := &pkgweapon.Weapon{}
	// if weapon.ID > 0 {
	// 	newWeapon = update(weapon)
	// } else {
	// 	newWeapon = insert(weapon)
	// }
	return newWeapon, nil
}

func (r *repository) FindById(id int64) (*pkgweapon.Weapon, error) {
	// weapons := generateWeapon()
	// for _, weapon := range weapons {
	// 	if weapon.ID == id {
	// 		return &weapon, nil
	// 	}
	// }
	return nil, fmt.Errorf("record not found")
}

func (r *repository) FindAll() (*[]pkgweapon.Weapon, error) {
	weapons := generateWeapon()
	return &weapons, nil
}

func (r *repository) Remove(id int64) error {
	// weapons := generateWeapon()
	// for index := range weapons {
	// 	if weapons[index].ID == id {
	// 		weapons = append(weapons[:index], weapons[index+1:]...)
	// 		break
	// 	}
	// }

	return nil
}

func insert(weapon *pkgweapon.Weapon) *pkgweapon.Weapon {
	// weapons := generateWeapon()
	// weapon.ID = int64(len(weapons) + 1)
	// weapons = append(weapons, *weapon)
	return nil
}

func update(weapon *pkgweapon.Weapon) *pkgweapon.Weapon {
	weapons := generateWeapon()
	if len(weapons) > 0 {
		for index := range weapons {
			if weapons[index].ID == weapon.ID {
				weapons[index].Description = weapon.Description
				weapons[index].Type = pkgweapon.TypeWeapon{
					Category:    weapon.Type.Category,
					Description: weapon.Type.Description,
				}
				return &weapons[index]

			}
		}
	}
	return nil
}

func generateWeapon() []pkgweapon.Weapon {
	weapons := make([]pkgweapon.Weapon, 0)
	// for i := 0; i < 10; i++ {

	// 	weapons = append(weapons, pkgweapon.Weapon{
	// 		ID:          int64(i),
	// 		Description: "Descrição da Arma",
	// 		Type: pkgweapon.TypeWeapon{
	// 			Category:    "arma branca",
	// 			Description: "descrição do tipo de arma",
	// 		},
	// 	})
	// }
	return weapons
}
