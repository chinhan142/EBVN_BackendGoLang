package service

import (
	"github.com/chinhan142/EBVN_BackendGoLang/internal/model"
	"github.com/google/uuid"
)

type HealthCheck interface {
	GetHealthCheck() model.Healthcheck
}

type healthCheckService struct {
	cfg *Config
}

func NewHealthCheck(cfg *Config) HealthCheck {
	if cfg.InstanceID == "" {
		cfg.InstanceID = uuid.New().String()
	}
	return &healthCheckService{
		cfg: cfg,
	}
}

func (s *healthCheckService) GetHealthCheck() model.Healthcheck {
	return model.Healthcheck{
		Message:     "OK",
		ServiceName: s.cfg.ServiceName,
		InstanceID:  s.cfg.InstanceID,
	}
}
