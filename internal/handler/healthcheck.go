package handler

import (
	"net/http"

	"github.com/chinhan142/EBVN_BackendGoLang/internal/service"
	"github.com/gin-gonic/gin"
)

// HealthCheck defines the interface for health check HTTP handler.
type HealthCheck interface {
	GetHealthCheck(c *gin.Context)
}

type healthCheckHandler struct {
	healthCheckService service.HealthCheck
}

// NewHealthCheck creates a new HealthCheck handler instance.
func NewHealthCheck(healthChecksvc service.HealthCheck) HealthCheck {
	return &healthCheckHandler{
		healthCheckService: healthChecksvc,
	}
}
func (s *healthCheckHandler) GetHealthCheck(c *gin.Context) {
	healthStatus := s.healthCheckService.GetHealthCheck()
	c.JSON(http.StatusOK, healthStatus)
}
