package weapon

import (
	pkgentityutils "delegacia.com.br/app/domain/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Weapon struct {
	ID          primitive.ObjectID         `bson:"_id,omitempty"`
	Description string                     `bson:"description"`
	Type        TypeWeapon                 `bson:"type"`
	EventRecord pkgentityutils.EventRecord `bson="eventRecord"`
}

type TypeWeapon struct {
	Category    string `bson:"category"`
	Description string `bson:"description"`
}
