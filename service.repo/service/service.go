package service

import (
	"fmt"

	"github.com/quangbuule/go-microservices-example/service.repo/store"

	"github.com/gin-gonic/gin"
	"github.com/quangbuule/go-microservices-example/service.repo/routes"
)

type Service struct {
	Name   string
	Port   string
	Engine *gin.Engine
	Store  *store.Engine
}

func New() *Service {
	return &Service{
		Name: "service.repo",
	}
}

func (s *Service) Initialize() {
	// Get options from process arguments
	c := getOptions()

	s.Port = c.Port

	// Initialize store
	s.Store = store.New()
	s.Store.InitializeState()

	// Initialize service engine
	s.Engine = gin.Default()

	s.Engine.Use(gin.Logger())
	s.Engine.Use(gin.Recovery())

	routes.AddRepoRoutes(s.Store, s.Engine)
}

func (s *Service) Run() {
	s.Engine.Run(fmt.Sprintf(":%s", s.Port))
}
