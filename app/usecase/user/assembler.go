package user_usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAssembler struct {
	ID       primitive.ObjectID `json:"oid,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type FindByEmailAndPassordAssembler struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAssembler struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
