package weapon

import (
	"strconv"

	pkgvictim "delegacia.com.br/app/domain/victim"
)

type repository struct{}

func NewVictimRepository() pkgvictim.VictimRepository {
	return &repository{}
}

func (r *repository) Save(victim pkgvictim.Victim) (*pkgvictim.Victim, error) {

	newVictim := &pkgvictim.Victim{}
	if victim.ID > 0 {
		newVictim = update(&victim)
	} else {
		newVictim = insert(&victim)
	}
	return newVictim, nil
}

func insert(victim *pkgvictim.Victim) *pkgvictim.Victim {
	victims := generateVictims()
	victim.ID = int64(len(victims) + 1)
	victims = append(victims, *victim)
	return victim
}

func update(victim *pkgvictim.Victim) *pkgvictim.Victim {
	victims := generateVictims()
	if len(victims) > 0 {
		for index := range victims {
			if victims[index].ID == victim.ID {
				victims[index].CPF = victim.CPF
				victims[index].Email = victim.Email
				victims[index].Name = victim.Name
				victims[index].Telephone = victim.Telephone
				return &victims[index]
			}
		}
	}
	return nil
}

func (r *repository) FindById(id int64) (*pkgvictim.Victim, error) {
	victims := generateVictims()
	return getVictim(victims, id), nil
}
func (r *repository) FindAll() (*[]pkgvictim.Victim, error) {
	victims := generateVictims()
	return &victims, nil
}
func (r *repository) Remove(id int64) error {
	return nil
}

func getVictim(victims []pkgvictim.Victim, id int64) *pkgvictim.Victim {
	if len(victims) > 0 {
		for _, victim := range victims {
			if victim.ID == id {
				return &victim
			}
		}
	}
	return nil
}

func generateVictims() []pkgvictim.Victim {
	victims := make([]pkgvictim.Victim, 0)
	for i := 0; i < 10; i++ {
		indiceStr := strconv.Itoa(i)
		victims = append(victims, pkgvictim.Victim{
			ID:        int64(i),
			Name:      "Vítima " + indiceStr,
			CPF:       "12345678" + indiceStr,
			Telephone: "(47) 9999-960" + indiceStr,
			Email:     "teste@teste.com.br",
		})
	}
	return victims
}
