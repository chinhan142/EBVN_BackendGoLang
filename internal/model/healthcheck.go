package model

type Healthcheck struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}
