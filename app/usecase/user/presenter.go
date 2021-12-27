package user_usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPresenter struct {
	ID    primitive.ObjectID `json:"oid,omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}
