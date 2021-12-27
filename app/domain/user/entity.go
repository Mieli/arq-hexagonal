package user

import (
	pkgentityutils "delegacia.com.br/app/domain/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID         `bson:"_id,omitempty"`
	Name        string                     `bson:"name"`
	Email       string                     `bson:"email"`
	Password    string                     `bson:"password"`
	EventRecord pkgentityutils.EventRecord `bson:"eventRecord"`
}
