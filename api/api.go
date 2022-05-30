package api

import (
	"fmt"
	"net/http"
	"privy-test/api/http/request"
	v1 "privy-test/api/v1"
	"privy-test/infra"
	"privy-test/manager"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server interface {
	Run()
}

type server struct {
	router             *echo.Echo
	requestParser      request.Parser
	serviceManager     manager.ServiceManager
	integrationManager manager.IntegrationManager
	infra              infra.Infra
}

func NewServer(infra infra.Infra) Server {
	return &server{
		router:             echo.New(),
		requestParser:      request.NewDefaultParser(),
		integrationManager: manager.NewIntegrationManager(infra),
		serviceManager:     manager.NewServiceManager(infra),
		infra:              infra,
	}
}

func (c *server) commonRouter(group *echo.Group) {
	healthHandler := v1.NewHealthCheckHandler(c.requestParser, c.serviceManager.HealthCheck())
	group.GET("/ping", healthHandler.HealthCheck)
	if c.infra.Config().Swagger.Enabled {
		group.GET("/swagger/*",
			echoSwagger.WrapHandler,
		)
	}
}

func (c *server) main(group *echo.Group) {
	cakeHandler := v1.NewCakeHandler(c.requestParser, c.integrationManager.Logger(), c.serviceManager.Cake())

	group.POST("/cake", cakeHandler.CreateCake)
	group.PATCH("/cake/:id", cakeHandler.UpdateCake)
	group.DELETE("/cake/:id", cakeHandler.DeleteCake)
	group.GET("/cake/:id", cakeHandler.GetDetailCake)
	group.GET("/cake", cakeHandler.GetList)
}

func (c *server) Handlers() {
	groupRouter := c.router.Group("api")
	c.commonRouter(groupRouter)
	c.main(groupRouter)
}

func (c *server) Run() {
	c.Handlers()
	c.run()
}

func (c *server) run() {
	addr := fmt.Sprintf("%s:%v", c.infra.Config().Host, c.infra.Config().Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      c.router,
		ReadTimeout:  time.Duration(c.infra.Config().ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.infra.Config().WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(c.infra.Config().IdleTimeout) * time.Second,
	}

	logger := c.infra.Logger().With().Str("component", "app.api").Logger()
	logger.Info(). // commit and build already set in infra.Logger
			Str("address", addr).
			Str("namespace", c.infra.Config().Namespace).
			Str("goVersion", c.infra.Config().GoVersion).
			Str("buildVersion", c.infra.Config().BuildVersion).
			Str("environment", c.infra.Config().Environment).
			Msg("server info")

	if err := c.router.StartServer(server); err != nil {
		logger.Fatal().Str("error", err.Error()).Msg("Server start error")
	}
}
