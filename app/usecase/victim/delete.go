package victim_usecase

import (
	"fmt"

	pkgvictim "delegacia.com.br/app/domain/victim"
)

type DeleteUseCase struct {
	VictimService pkgvictim.Service
	ID            *int64
}

type DeleteUseCaseParams struct {
	VictimService pkgvictim.Service
}

func NewDeleteUseCase(params DeleteUseCaseParams) DeleteUseCase {
	return DeleteUseCase{
		VictimService: params.VictimService,
	}
}

func (uc *DeleteUseCase) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.VictimService.Remove(*uc.ID)
	if err != nil {
		return fmt.Errorf("error remove data")
	}
	return nil
}
