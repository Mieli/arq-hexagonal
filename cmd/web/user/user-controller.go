package user_controller

import (
	"fmt"
	"net/http"

	pkguseruc "delegacia.com.br/app/usecase/user"
	"github.com/labstack/echo"
)

type UserController struct {
	InsertUseCaseParams   pkguseruc.InsertUseCaseParams
	FindAllUseCaseParams  pkguseruc.FindAllUseCaseParams
	FindByIdUseCaseParams pkguseruc.FindByIDUseCaseParams
	UpdateUseCaseParams   pkguseruc.UpdateUseCaseParams
	DeleteUseCaseParams   pkguseruc.DeleteUseCaseParams
}

type UserControllerParams struct {
	InsertUseCaseParams   pkguseruc.InsertUseCaseParams
	FindAllUseCaseParams  pkguseruc.FindAllUseCaseParams
	FindByIdUseCaseParams pkguseruc.FindByIDUseCaseParams
	UpdateUseCaseParams   pkguseruc.UpdateUseCaseParams
	DeleteUseCaseParams   pkguseruc.DeleteUseCaseParams
}

func NewUserController(params UserControllerParams, g *echo.Group) {
	controller := UserController{
		InsertUseCaseParams:   params.InsertUseCaseParams,
		FindAllUseCaseParams:  params.FindAllUseCaseParams,
		FindByIdUseCaseParams: params.FindByIdUseCaseParams,
		UpdateUseCaseParams:   params.UpdateUseCaseParams,
		DeleteUseCaseParams:   params.DeleteUseCaseParams,
	}

	g.POST("/user", controller.Insert)
	g.GET("/users", controller.FindAll)
	g.GET("/user/:id", controller.FindById)
	g.PUT("/user/:id", controller.Update)
	g.DELETE("/user/:id", controller.Remove)

}

func (c *UserController) Insert(ctx echo.Context) error {
	assembler := pkguseruc.UserAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkguseruc.NewInsertUserUseCase(c.InsertUseCaseParams)
	uc.Assembler = &assembler
	user, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	return ctx.JSON(http.StatusOK, user)

}

func (c *UserController) Update(ctx echo.Context) error {

	assembler := pkguseruc.UserAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkguseruc.NewUpdateUseCase(c.UpdateUseCaseParams)
	uc.Assembler = &assembler
	victim, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	return ctx.JSON(http.StatusOK, victim)
}

func (c *UserController) FindAll(ctx echo.Context) error {

	uc := pkguseruc.NewFindAllUseCase(c.FindAllUseCaseParams)
	data, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	return ctx.JSON(http.StatusOK, data)

}

func (c *UserController) FindById(ctx echo.Context) error {

	id := ctx.Param("id")
	if id == "" {
		return fmt.Errorf("Invalid id")
	}
	uc := pkguseruc.NewFindByIdUseCase(c.FindByIdUseCaseParams)
	uc.ID = id
	data, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *UserController) Remove(ctx echo.Context) error {

	id := ctx.Param("id")
	if id == "" {
		return fmt.Errorf("Invalid id")
	}
	uc := pkguseruc.NewDeleteUseCase(c.DeleteUseCaseParams)
	uc.ID = &id
	err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, nil)
}
