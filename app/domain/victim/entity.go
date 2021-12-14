package victim

import (
	pkgentityutils "delegacia.com.br/app/domain/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Victim struct {
	ID          primitive.ObjectID         `bson:"_id,omitempty"`
	Name        string                     `bson:"name"`
	CPF         string                     `bson:"cpf"`
	Telephone   string                     `bson:"telephone"`
	Email       string                     `bson:"email"`
	EventRecord pkgentityutils.EventRecord `bson:"eventRecord"`
}
