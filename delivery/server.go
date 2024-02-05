package delivery

import (
	"fmt"
	"log"

	"github.com/eulbyvan/enigma-university/config"
	"github.com/eulbyvan/enigma-university/delivery/controller"
	"github.com/eulbyvan/enigma-university/manager"
	"github.com/gin-gonic/gin"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (s *Server) setupControllers() { // ERROR DI SEKITAR SINI
	rg := s.engine.Group("/api/v1")
	controller.NewUserController(s.useCaseManager.NewUserUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	inframanager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(inframanager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	log.Println("DISINIIIIIIIIIIII>>>>>>>>>", engine)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
	}
}
