package weapon_usecase

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeaponPresenter struct {
	ID          primitive.ObjectID  `json:"oid,omitempty"`
	Description string              `json:"description"`
	Type        TypeWeaponPresenter `json:"type"`
}

type TypeWeaponPresenter struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}
