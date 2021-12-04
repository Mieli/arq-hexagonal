package victim

type Victim struct {
	ID        int64  `bson:"_id"`
	Name      string `bson:"name"`
	CPF       string `bson:"cpf"`
	Telephone string `bson:"telephone"`
	Email     string `bson:"email"`
}
