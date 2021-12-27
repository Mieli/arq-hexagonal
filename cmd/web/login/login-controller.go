package login_controller

import (
	"net/http"

	pkguseruc "delegacia.com.br/app/usecase/user"
	pkgservices "delegacia.com.br/infra/services"
	"github.com/labstack/echo"
)

type LoginController struct {
	LoginUseCaseParams pkguseruc.LoginUseCaseParams
}
type LoginControllerParams struct {
	LoginUseCaseParams pkguseruc.LoginUseCaseParams
}

func NewLoginController(params LoginControllerParams, g *echo.Group) {
	controller := LoginController{
		LoginUseCaseParams: params.LoginUseCaseParams,
	}

	g.POST("/login", controller.Login)

}

func (c LoginController) Login(ctx echo.Context) error {
	assembler := pkguseruc.LoginAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	uc := pkguseruc.NewLoginUseCase(c.LoginUseCaseParams)
	uc.Assembler = &assembler
	data, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	jwt := pkgservices.NewJwtService()
	token, err := jwt.GenerateToken(data.ID.Hex())

	return ctx.JSON(http.StatusOK, token)
}
