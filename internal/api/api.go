package api

import (
	"fmt"

	"github.com/chinhan142/EBVN_BackendGoLang/internal/handler"
	"github.com/chinhan142/EBVN_BackendGoLang/internal/service"
	"github.com/gin-gonic/gin"
)

// Engine defines the HTTP server engine interface.
type Engine interface {
	Start() error
}

type engine struct {
	app *gin.Engine
	cfg *Config
}

// NewEngine creates a new Engine instance and registers all routes.
func NewEngine(cfg *Config) Engine {
	app := &engine{
		app: gin.Default(),
		cfg: cfg,
	}

	app.initRoutes()

	return app
}

func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

func (e *engine) initRoutes() {
	healthCheckSvc := service.NewHealthCheck(e.cfg.ServiceName, e.cfg.InstanceID)
	healthCheckHandler := handler.NewHealthCheck(healthCheckSvc)

	e.app.GET("/health-check", healthCheckHandler.GetHealthCheck)
}
