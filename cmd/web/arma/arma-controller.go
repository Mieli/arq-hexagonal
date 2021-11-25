package arma

import (
	"net/http"
	"strconv"

	pkgarmauc "delegacia.com.br/app/usecase/arma"
	"github.com/labstack/echo"
)

type ArmaController struct {
	SaveUseCaseParams     pkgarmauc.SaveUseCaseParams
	FindAllUseCaseParams  pkgarmauc.FindAllUseCaseParams
	FindByIdUseCaseParams pkgarmauc.FindByIdUseCaseParams
	DeleteUseCaseParams   pkgarmauc.DeleteUseCaseParams
}

type ArmaControlleParams struct {
	SaveUseCaseParams     pkgarmauc.SaveUseCaseParams
	FindAllUseCaseParams  pkgarmauc.FindAllUseCaseParams
	FindByIdUseCaseParams pkgarmauc.FindByIdUseCaseParams
	DeleteUseCaseParams   pkgarmauc.DeleteUseCaseParams
}

func NewArmaController(params *ArmaControlleParams, g *echo.Group) {
	controller := ArmaController{
		SaveUseCaseParams:     params.SaveUseCaseParams,
		FindAllUseCaseParams:  params.FindAllUseCaseParams,
		FindByIdUseCaseParams: params.FindByIdUseCaseParams,
		DeleteUseCaseParams:   params.DeleteUseCaseParams,
	}
	g.POST("/arma", controller.Save)
	g.GET("/armas", controller.FindAll)
	g.GET("/arma/:id", controller.FindById)
	g.DELETE("/arma/:id", controller.Remove)

}

func (c *ArmaController) Save(ctx echo.Context) error {

	assembler := pkgarmauc.ArmaAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgarmauc.NewSaveArmaUseCase(c.SaveUseCaseParams)
	uc.Assembler = &assembler
	arma, err := uc.Eecute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, arma)

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
