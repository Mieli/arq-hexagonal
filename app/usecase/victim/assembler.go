package victim_usecase

type VictimAssembler struct {
	ID        int64  `json:"_id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
}
