package arma

type ArmaPresenter struct {
	ID          int64             `json:"oid"`
	Description string            `json:"description"`
	Type        TypeArmaPresenter `json:"type"`
}

type TypeArmaPresenter struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}
