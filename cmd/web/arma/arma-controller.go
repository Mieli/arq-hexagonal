package arma

import (
	"fmt"
	"net/http"
	"strconv"

	pkgarmauc "delegacia.com.br/app/usecase/arma"
	"github.com/labstack/echo"
)

type ArmaController struct {
	FindAllUseCaseParams  pkgarmauc.FindAllUseCaseParams
	FindByIdUseCaseParams pkgarmauc.FindByIdUseCaseParams
	InsertUseCaseParams   pkgarmauc.InsertUseCaseParams
	UpdateUseCaseParams   pkgarmauc.UpdateUseCaseParams
	DeleteUseCaseParams   pkgarmauc.DeleteUseCaseParams
}

type ArmaControlleParams struct {
	FindAllUseCaseParams  pkgarmauc.FindAllUseCaseParams
	FindByIdUseCaseParams pkgarmauc.FindByIdUseCaseParams
	InsertUseCaseParams   pkgarmauc.InsertUseCaseParams
	UpdateUseCaseParams   pkgarmauc.UpdateUseCaseParams
	DeleteUseCaseParams   pkgarmauc.DeleteUseCaseParams
}

func NewArmaController(params *ArmaControlleParams, g *echo.Group) {
	controller := ArmaController{
		FindAllUseCaseParams:  params.FindAllUseCaseParams,
		FindByIdUseCaseParams: params.FindByIdUseCaseParams,
		InsertUseCaseParams:   params.InsertUseCaseParams,
		UpdateUseCaseParams:   params.UpdateUseCaseParams,
		DeleteUseCaseParams:   params.DeleteUseCaseParams,
	}
	g.POST("/arma", controller.Insert)
	g.GET("/armas", controller.FindAll)
	g.GET("/arma/:id", controller.FindById)
	g.PUT("/arma/:id", controller.Update)
	g.DELETE("/arma/:id", controller.Remove)

}

func (c *ArmaController) Insert(ctx echo.Context) error {

	assembler := pkgarmauc.ArmaAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgarmauc.NewInsertArmaUseCase(c.InsertUseCaseParams)
	uc.Assembler = &assembler
	arma, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, arma)

}
func (c *ArmaController) Update(ctx echo.Context) error {

	idString := ctx.Param("id")
	if id, err := strconv.ParseInt(idString, 10, 64); err == nil {

		assembler := pkgarmauc.ArmaAssembler{}
		if err := ctx.Bind(&assembler); err != nil {
			return ctx.JSON(http.StatusPreconditionFailed, err)
		}

		if id != assembler.ID {
			return fmt.Errorf("id invalid")
		}

		uc := pkgarmauc.NewUpdateUseCase(c.UpdateUseCaseParams)
		uc.Assembler = &assembler
		arma, err := uc.Execute()
		if err != nil {
			return ctx.JSON(http.StatusPreconditionFailed, nil)
		}
		return ctx.JSON(http.StatusOK, arma)
	}
	return ctx.JSON(http.StatusPreconditionFailed, nil)
}

func (c *ArmaController) FindAll(ctx echo.Context) error {

	uc := pkgarmauc.NewFindAllUseCase(c.FindAllUseCaseParams)
	armas, err := uc.Execute()

	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, armas)

}

func (c *ArmaController) FindById(ctx echo.Context) error {

	idAssembler := ctx.Param("id")
	id, err := strconv.ParseInt(idAssembler, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgarmauc.NewFindByIdUseCase(c.FindByIdUseCaseParams)
	uc.ID = &id
	arma, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, arma)
}

func (c *ArmaController) Remove(ctx echo.Context) error {

	idAssembler := ctx.Param("id")
	id, err := strconv.ParseInt(idAssembler, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgarmauc.NewDeleteUseCase(c.DeleteUseCaseParams)
	uc.ID = &id
	err = uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, nil)
}
