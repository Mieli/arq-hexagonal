package victim_usecase

import (
	pkgvictim "delegacia.com.br/app/domain/victim"
)

func GenerateVictim(assembler *VictimAssembler) *pkgvictim.Victim {
	if assembler != nil {
		return &pkgvictim.Victim{
			ID:        assembler.ID,
			Name:      assembler.Name,
			CPF:       assembler.CPF,
			Telephone: assembler.Telephone,
			Email:     assembler.Email,
		}
	}
	return nil
}
func GenerateVictimPresenter(victim *pkgvictim.Victim) *VictimPresenter {
	if victim != nil {
		return &VictimPresenter{
			ID:        victim.ID,
			Name:      victim.Name,
			CPF:       victim.CPF,
			Telephone: victim.Telephone,
			Email:     victim.Email,
		}
	}
	return nil
}

func GenerateListPresenter(victims []pkgvictim.Victim) []*VictimPresenter {
	list := make([]*VictimPresenter, 0)
	if len(victims) > 0 {
		for _, victim := range victims {
			list = append(list, &VictimPresenter{
				ID:        victim.ID,
				Name:      victim.Name,
				CPF:       victim.CPF,
				Telephone: victim.Telephone,
				Email:     victim.Email,
			})
		}
	}
	return list
}
