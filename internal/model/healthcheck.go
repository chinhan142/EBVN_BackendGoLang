package model

// Healthcheck represents the health check response payload.
type Healthcheck struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}
