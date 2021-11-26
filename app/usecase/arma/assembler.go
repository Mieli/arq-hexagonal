package arma

type ArmaAssembler struct {
	ID          int64             `json:"id"`
	Description string            `json:"description"`
	Type        TypeArmaAssembler `json:"type"`
}

type TypeArmaAssembler struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}
