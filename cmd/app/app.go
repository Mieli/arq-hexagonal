package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pkgarma "delegacia.com.br/app/domain/arma"
	pkgarmauc "delegacia.com.br/app/usecase/arma"
	pkgarmacontroller "delegacia.com.br/cmd/web/arma"
	pkgarmainfra "delegacia.com.br/infra/database/repositories/arma"
	"github.com/labstack/echo"
)

type Server struct {
	server *http.Server
}

type dependenceParams struct {
	ArmaService pkgarma.Service
}

func buildDependeciesParams() dependenceParams {
	params := dependenceParams{}
	params.ArmaService = pkgarma.NewServiceArma(pkgarmainfra.NewArmaRepository())

	return params
}

func buildArmaEndPoints(dependency *dependenceParams, g *echo.Group) {

	findAllParams := pkgarmauc.FindAllUseCaseParams{
		Service: dependency.ArmaService,
	}

	findByIdParams := pkgarmauc.FindByIdUseCaseParams{
		Service: dependency.ArmaService,
	}

	insertParams := pkgarmauc.InsertUseCaseParams{
		Service: dependency.ArmaService,
	}

	updateParams := pkgarmauc.UpdateUseCaseParams{
		Service: dependency.ArmaService,
	}

	deleteParams := pkgarmauc.DeleteUseCaseParams{
		Service: dependency.ArmaService,
	}

	armaControlleParams := pkgarmacontroller.ArmaControlleParams{
		FindAllUseCaseParams:  findAllParams,
		FindByIdUseCaseParams: findByIdParams,
		InsertUseCaseParams:   insertParams,
		UpdateUseCaseParams:   updateParams,
		DeleteUseCaseParams:   deleteParams,
	}
	pkgarmacontroller.NewArmaController(&armaControlleParams, g)
}

func Start() {

	router := echo.New()
	routerGroup := router.Group("/api/v1")

	dependency := buildDependeciesParams()
	buildArmaEndPoints(&dependency, routerGroup)

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
