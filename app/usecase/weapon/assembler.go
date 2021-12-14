package weapon_usecase

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeaponAssembler struct {
	ID          primitive.ObjectID  `json:"id"`
	Description string              `json:"description"`
	Type        TypeWeaponAssembler `json:"type"`
}

type TypeWeaponAssembler struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}
