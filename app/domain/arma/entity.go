package arma

type Arma struct {
	ID          int64    `bson:"_id"`
	Description string   `bson:"description"`
	Type        TypeArma `bson:"type"`
}

type TypeArma struct {
	Category    string `bson:"category"`
	Description string `bson:"description"`
}
