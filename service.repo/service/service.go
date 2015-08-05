package service

import (
	"fmt"
	"time"

	"github.com/agencyrevolution/go-microservices-example/service.repo/routes"
	"github.com/agencyrevolution/go-microservices-example/service.repo/store"
	"github.com/agencyrevolution/go-microservices-example/utils"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Name          string
	Port          string
	Router        *gin.Engine
	Store         *store.Engine
	VulcandClient *utils.VulcandClient
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

	// Initialize vulcan client
	s.VulcandClient = utils.NewVulcanClient(s.Name, s.Port, 10*time.Second)

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

func (s *Service) KeepAlive() {
	s.VulcandClient.KeepAlive()
}
