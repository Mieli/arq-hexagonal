package weaponcontroller

import (
	"net/http"
	"strconv"

	pkgweaponuc "delegacia.com.br/app/usecase/weapon"
	"github.com/labstack/echo"
)

type WeaponController struct {
	FindAllUseCaseParams  pkgweaponuc.FindAllUseCaseParams
	FindByIdUseCaseParams pkgweaponuc.FindByIdUseCaseParams
	InsertUseCaseParams   pkgweaponuc.InsertUseCaseParams
	UpdateUseCaseParams   pkgweaponuc.UpdateUseCaseParams
	DeleteUseCaseParams   pkgweaponuc.DeleteUseCaseParams
}

type WeaponControlleParams struct {
	FindAllUseCaseParams  pkgweaponuc.FindAllUseCaseParams
	FindByIdUseCaseParams pkgweaponuc.FindByIdUseCaseParams
	InsertUseCaseParams   pkgweaponuc.InsertUseCaseParams
	UpdateUseCaseParams   pkgweaponuc.UpdateUseCaseParams
	DeleteUseCaseParams   pkgweaponuc.DeleteUseCaseParams
}

func NewWeaponController(params *WeaponControlleParams, g *echo.Group) {
	controller := WeaponController{
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

func (c *WeaponController) Insert(ctx echo.Context) error {

	assembler := pkgweaponuc.WeaponAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgweaponuc.NewInsertUseCase(c.InsertUseCaseParams)
	uc.Assembler = &assembler
	weapon, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, weapon)

}
func (c *WeaponController) Update(ctx echo.Context) error {

	// idString := ctx.Param("id")
	// if id, err := strconv.ParseInt(idString, 10, 64); err == nil {

	// 	assembler := pkgweaponuc.WeaponAssembler{}
	// 	if err := ctx.Bind(&assembler); err != nil {
	// 		return ctx.JSON(http.StatusPreconditionFailed, err)
	// 	}

	// 	if id != assembler.ID {
	// 		return fmt.Errorf("id invalid")
	// 	}

	// 	uc := pkgweaponuc.NewUpdateUseCase(c.UpdateUseCaseParams)
	// 	uc.Assembler = &assembler
	// 	weapon, err := uc.Execute()
	// 	if err != nil {
	// 		return ctx.JSON(http.StatusPreconditionFailed, nil)
	// 	}
	// 	return ctx.JSON(http.StatusOK, weapon)
	// }
	return ctx.JSON(http.StatusPreconditionFailed, nil)
}

func (c *WeaponController) FindAll(ctx echo.Context) error {

	uc := pkgweaponuc.NewFindAllUseCase(c.FindAllUseCaseParams)
	weapons, err := uc.Execute()

	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, weapons)

}

func (c *WeaponController) FindById(ctx echo.Context) error {

	idAssembler := ctx.Param("id")
	id, err := strconv.ParseInt(idAssembler, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgweaponuc.NewFindByIdUseCase(c.FindByIdUseCaseParams)
	uc.ID = &id
	weapon, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, weapon)
}

func (c *WeaponController) Remove(ctx echo.Context) error {

	idAssembler := ctx.Param("id")
	id, err := strconv.ParseInt(idAssembler, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkgweaponuc.NewDeleteUseCase(c.DeleteUseCaseParams)
	uc.ID = &id
	err = uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, nil)
}
