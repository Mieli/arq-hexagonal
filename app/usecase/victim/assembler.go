package victim_usecase

import "go.mongodb.org/mongo-driver/bson/primitive"

type VictimAssembler struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	CPF       string             `json:"cpf"`
	Telephone string             `json:"telephone"`
	Email     string             `json:"email"`
}
