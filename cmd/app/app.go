package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pkguser "delegacia.com.br/app/domain/user"
	pkgvictim "delegacia.com.br/app/domain/victim"
	pkgweapon "delegacia.com.br/app/domain/weapon"
	pkguseruc "delegacia.com.br/app/usecase/user"
	pkgvictimuc "delegacia.com.br/app/usecase/victim"
	pkgweaponuc "delegacia.com.br/app/usecase/weapon"
	pkglogincontroller "delegacia.com.br/cmd/web/login"
	pkgusercontroller "delegacia.com.br/cmd/web/user"
	pkgvictimcontroller "delegacia.com.br/cmd/web/victim"
	pkgweaponacontroller "delegacia.com.br/cmd/web/weapon"
	pkguserinfra "delegacia.com.br/infra/database/repositories/user"
	pkgvictiminfra "delegacia.com.br/infra/database/repositories/victim"
	pkgweaponinfra "delegacia.com.br/infra/database/repositories/weapon"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	server *http.Server
}

type dependenceParams struct {
	WeaponService pkgweapon.Service
	VictimService pkgvictim.Service
	UserService   pkguser.Service
}

func buildDependeciesParams(db mongo.Client) dependenceParams {

	params := dependenceParams{}

	params.WeaponService = pkgweapon.NewServiceWeapon(pkgweaponinfra.NewWeaponRepository(db))
	params.VictimService = pkgvictim.NewServiceVictim(pkgvictiminfra.NewVictimRepository(db))
	params.UserService = pkguser.NewServiceUser(pkguserinfra.NewUserRepository(db))
	return params
}

func buildWeaponEndPoint(dependency *dependenceParams, g *echo.Group) {

	findAllParams := pkgweaponuc.FindAllUseCaseParams{
		Service: dependency.WeaponService,
	}

	findByIdParams := pkgweaponuc.FindByIdUseCaseParams{
		Service: dependency.WeaponService,
	}

	insertParams := pkgweaponuc.InsertUseCaseParams{
		Service: dependency.WeaponService,
	}

	updateParams := pkgweaponuc.UpdateUseCaseParams{
		Service: dependency.WeaponService,
	}

	deleteParams := pkgweaponuc.DeleteUseCaseParams{
		Service: dependency.WeaponService,
	}

	weaponControlleParams := pkgweaponacontroller.WeaponControlleParams{
		FindAllUseCaseParams:  findAllParams,
		FindByIdUseCaseParams: findByIdParams,
		InsertUseCaseParams:   insertParams,
		UpdateUseCaseParams:   updateParams,
		DeleteUseCaseParams:   deleteParams,
	}
	pkgweaponacontroller.NewWeaponController(&weaponControlleParams, g)
}

func buildVictimEndPoints(dependency *dependenceParams, g *echo.Group) {

	findAllParamns := pkgvictimuc.FindAllUseCaseParams{
		VictimService: dependency.VictimService,
	}

	findByIdParams := pkgvictimuc.FindByIdUseCaseParams{
		VictimService: dependency.VictimService,
	}
	deleteParams := pkgvictimuc.DeleteUseCaseParams{
		VictimService: dependency.VictimService,
	}
	insertParams := pkgvictimuc.InsertUseCaseParams{
		VictimService: dependency.VictimService,
	}

	updateParams := pkgvictimuc.UpdateUseCaseParams{
		VictimService: dependency.VictimService,
	}

	victimControllerParams := pkgvictimcontroller.VictimControllerParams{
		FindAllUseCaseParams:  findAllParamns,
		FindByIdUseCaseParams: findByIdParams,
		DeleteUseCaseParams:   deleteParams,
		InsertUseCaseParams:   insertParams,
		UpdateUseCaseParams:   updateParams,
	}

	pkgvictimcontroller.NewVictimController(victimControllerParams, g)
}

func buildUserEndPoints(dependency *dependenceParams, g *echo.Group) {

	findAllParamns := pkguseruc.FindAllUseCaseParams{
		Service: dependency.UserService,
	}

	findByIdParams := pkguseruc.FindByIDUseCaseParams{
		Service: dependency.UserService,
	}
	deleteParams := pkguseruc.DeleteUseCaseParams{
		Service: dependency.UserService,
	}
	insertParams := pkguseruc.InsertUseCaseParams{
		Service: dependency.UserService,
	}

	updateParams := pkguseruc.UpdateUseCaseParams{
		Service: dependency.UserService,
	}

	userControllerParams := pkgusercontroller.UserControllerParams{
		InsertUseCaseParams:   insertParams,
		FindAllUseCaseParams:  findAllParamns,
		FindByIdUseCaseParams: findByIdParams,
		UpdateUseCaseParams:   updateParams,
		DeleteUseCaseParams:   deleteParams,
	}

	pkgusercontroller.NewUserController(userControllerParams, g)
}

func buildLoginEndPoints(dependency *dependenceParams, g *echo.Group) {

	loginParams := pkguseruc.LoginUseCaseParams{
		Service: dependency.UserService,
	}

	loginControllerParams := pkglogincontroller.LoginControllerParams{
		LoginUseCaseParams: loginParams,
	}

	pkglogincontroller.NewLoginController(loginControllerParams, g)
}

func Start(db *mongo.Client) {

	router := echo.New()
	routerGroup := router.Group("/api/v1")

	dependency := buildDependeciesParams(*db)
	buildWeaponEndPoint(&dependency, routerGroup)
	buildVictimEndPoints(&dependency, routerGroup)
	buildUserEndPoints(&dependency, routerGroup)
	buildLoginEndPoints(&dependency, routerGroup)

	server := newServer("6000", router)
	server.ListenAndServe()

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
}

func newServer(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * 60 * time.Second,
			WriteTimeout: 5 * 60 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() {
	go func() {
		fmt.Println("Server runing in port 6000")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("erro: %s", err)
		}
	}()

}
