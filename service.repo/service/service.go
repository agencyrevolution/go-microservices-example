package service

import (
	"fmt"

	"github.com/agencyrevolution/go-microservices-example/service.repo/store"

	"github.com/agencyrevolution/go-microservices-example/service.repo/routes"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Name   string
	Port   string
	Router *gin.Engine
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
	s.Router = gin.Default()

	s.Router.Use(gin.Logger())
	s.Router.Use(gin.Recovery())

	routes.AddRepoRoutes(s.Store, s.Router)
}

func (s *Service) Run() {
	s.Router.Run(fmt.Sprintf(":%s", s.Port))
}
