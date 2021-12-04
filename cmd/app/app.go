package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pkgvictim "delegacia.com.br/app/domain/victim"
	pkgweapon "delegacia.com.br/app/domain/weapon"
	pkgvictimuc "delegacia.com.br/app/usecase/victim"
	pkgweaponuc "delegacia.com.br/app/usecase/weapon"
	pkgvictimcontroller "delegacia.com.br/cmd/web/victim"
	pkgweaponacontroller "delegacia.com.br/cmd/web/weapon"
	pkgvictiminfra "delegacia.com.br/infra/database/repositories/victim"
	pkgweaponinfra "delegacia.com.br/infra/database/repositories/weapon"
	"github.com/labstack/echo"
)

type Server struct {
	server *http.Server
}

type dependenceParams struct {
	WeaponService pkgweapon.Service
	VictimService pkgvictim.Service
}

func buildDependeciesParams() dependenceParams {
	params := dependenceParams{}

	params.WeaponService = pkgweapon.NewServiceWeapon(pkgweaponinfra.NewWeaponRepository())
	params.VictimService = pkgvictim.NewServiceVictim(pkgvictiminfra.NewVictimRepository())

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

func Start() {

	router := echo.New()
	routerGroup := router.Group("/api/v1")

	dependency := buildDependeciesParams()
	buildWeaponEndPoint(&dependency, routerGroup)
	buildVictimEndPoints(&dependency, routerGroup)

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
