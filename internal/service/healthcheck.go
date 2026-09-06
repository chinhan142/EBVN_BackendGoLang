package service

import (
	"uuid"

	"github.com/chinhan142/EBVN_BackendGoLang/internal/model"
)

type HealthCheck interface {
	GetHealthCheck() model.Healthcheck
}

type healthCheckService struct {
	serviceName string
	instanceID  string
}

func NewHealthCheck(serviceName string, instanceID string) HealthCheck {
	if instanceID == "" {
		instanceID = uuid.New().String()
	}

	return &healthCheckService{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}

func (s *healthCheckService) GetHealthCheck() model.Healthcheck {
	return model.Healthcheck{
		Message:     "OK",
		ServiceName: s.serviceName,
		InstanceID:  s.instanceID,
	}
}
