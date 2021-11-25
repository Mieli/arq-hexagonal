package arma

import "fmt"

type repository struct{}

func (r *repository) Save(arma *Arma) (*Arma, error) {
	armas := generateArmas()

	arma.ID = int64(len(armas) + 1)

	armas = append(armas, *arma)
	return arma, nil
}

func (r *repository) FindById(id int64) (*Arma, error) {
	armas := generateArmas()
	for _, arma := range armas {
		if arma.ID == id {
			return &arma, nil
		}
	}
	return nil, fmt.Errorf("record not found")
}

func (r *repository) FindAll() (*[]Arma, error) {
	armas := generateArmas()
	return &armas, nil
}

func (r *repository) Remove(id int64) error {
	armas := generateArmas()
	for index := range armas {
		if armas[index].ID == id {
			armas = append(armas[:index], armas[index+1:]...)
			break
		}
	}

	return nil
}

func generateArmas() []Arma {
	armas := make([]Arma, 0)
	for i := 0; i < 10; i++ {

		armas = append(armas, Arma{
			ID:          int64(i),
			Description: "Descrição da Arma",
			Type: TypeArma{
				Category:    "arma branca",
				Description: "descrição do tipo de arma",
			},
		})
	}
	return armas
}

func NewArmaRepository() ArmaRepository {
	return &repository{}
}
