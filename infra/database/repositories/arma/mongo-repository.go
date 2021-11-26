package arma

import (
	"fmt"

	pkgarma "delegacia.com.br/app/domain/arma"
)

type repository struct{}

func (r *repository) Save(arma *pkgarma.Arma) (*pkgarma.Arma, error) {
	newArma := &pkgarma.Arma{}
	if arma.ID > 0 {
		newArma = update(arma)
	} else {
		newArma = insert(arma)
	}
	return newArma, nil
}

func (r *repository) FindById(id int64) (*pkgarma.Arma, error) {
	armas := generateArmas()
	for _, arma := range armas {
		if arma.ID == id {
			return &arma, nil
		}
	}
	return nil, fmt.Errorf("record not found")
}

func (r *repository) FindAll() (*[]pkgarma.Arma, error) {
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

func insert(arma *pkgarma.Arma) *pkgarma.Arma {
	armas := generateArmas()
	arma.ID = int64(len(armas) + 1)
	armas = append(armas, *arma)
	return arma
}

func update(arma *pkgarma.Arma) *pkgarma.Arma {
	armas := generateArmas()
	if len(armas) > 0 {
		for index := range armas {
			if armas[index].ID == arma.ID {
				armas[index].Description = arma.Description
				armas[index].Type = pkgarma.TypeArma{
					Category:    arma.Type.Category,
					Description: arma.Type.Description,
				}
				return &armas[index]

			}
		}
	}
	return nil
}

func generateArmas() []pkgarma.Arma {
	armas := make([]pkgarma.Arma, 0)
	for i := 0; i < 10; i++ {

		armas = append(armas, pkgarma.Arma{
			ID:          int64(i),
			Description: "Descrição da Arma",
			Type: pkgarma.TypeArma{
				Category:    "arma branca",
				Description: "descrição do tipo de arma",
			},
		})
	}
	return armas
}

func NewArmaRepository() pkgarma.ArmaRepository {
	return &repository{}
}
