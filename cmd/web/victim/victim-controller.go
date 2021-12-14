package victim

import (
	"fmt"
	"net/http"

	pkgvictimuc "delegacia.com.br/app/usecase/victim"

	"github.com/labstack/echo"
)

type VictimController struct {
	FindAllUseCaseParams  pkgvictimuc.FindAllUseCaseParams
	FindByIdUseCaseParams pkgvictimuc.FindByIdUseCaseParams
	DeleteUseCaseParams   pkgvictimuc.DeleteUseCaseParams
	InsertUseCaseParams   pkgvictimuc.InsertUseCaseParams
	UpdateUseCaseParams   pkgvictimuc.UpdateUseCaseParams
}
type VictimControllerParams struct {
	FindAllUseCaseParams  pkgvictimuc.FindAllUseCaseParams
	FindByIdUseCaseParams pkgvictimuc.FindByIdUseCaseParams
	DeleteUseCaseParams   pkgvictimuc.DeleteUseCaseParams
	InsertUseCaseParams   pkgvictimuc.InsertUseCaseParams
	UpdateUseCaseParams   pkgvictimuc.UpdateUseCaseParams
}

func NewVictimController(params VictimControllerParams, g *echo.Group) {
	controller := VictimController{
		FindAllUseCaseParams:  params.FindAllUseCaseParams,
		FindByIdUseCaseParams: params.FindByIdUseCaseParams,
		DeleteUseCaseParams:   params.DeleteUseCaseParams,
		InsertUseCaseParams:   params.InsertUseCaseParams,
		UpdateUseCaseParams:   params.UpdateUseCaseParams,
	}

	g.POST("/vitima", controller.Insert)
	g.GET("/vitimas", controller.FindAll)
	g.GET("/vitima/:id", controller.FindById)
	g.PUT("/vitima/:id", controller.Update)
	g.DELETE("/vitima/:id", controller.Remove)

}

func (c *VictimController) Insert(ctx echo.Context) error {

	assembler := pkgvictimuc.VictimAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgvictimuc.NewInsertUseCase(c.InsertUseCaseParams)
	uc.Assembler = &assembler
	victim, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	return ctx.JSON(http.StatusOK, victim)
}

func (c *VictimController) Update(ctx echo.Context) error {

	assembler := pkgvictimuc.VictimAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgvictimuc.NewUpdateUseCase(c.UpdateUseCaseParams)
	uc.Assembler = &assembler
	victim, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	return ctx.JSON(http.StatusOK, victim)
}

func (c *VictimController) FindAll(ctx echo.Context) error {

	uc := pkgvictimuc.NewFindAllUseCase(c.FindAllUseCaseParams)
	data, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	return ctx.JSON(http.StatusOK, data)

}

func (c *VictimController) FindById(ctx echo.Context) error {

	id := ctx.Param("id")
	if id == "" {
		return fmt.Errorf("Invalid id")
	}
	uc := pkgvictimuc.NewFindByIdUseCase(c.FindByIdUseCaseParams)
	uc.ID = &id
	data, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *VictimController) Remove(ctx echo.Context) error {

	id := ctx.Param("id")
	if id == "" {
		return fmt.Errorf("Invalid id")
	}
	uc := pkgvictimuc.NewDeleteUseCase(c.DeleteUseCaseParams)
	uc.ID = &id
	err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, nil)
}
