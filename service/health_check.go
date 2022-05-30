package service

import (
	"context"
	"privy-test/infra"
	"privy-test/param/healthcheck"
)

type healthCheckService struct {
	config infra.MergeConfig
}

func NewHealthCheckService(config infra.MergeConfig) HealthCheckService {
	return &healthCheckService{
		config: config,
	}
}

func (h healthCheckService) HealthCheck(ctx context.Context) healthcheck.HTTPHealthCheckResponse {
	return healthcheck.HTTPHealthCheckResponse{
		Data: healthcheck.DataResponse{
			Environment: h.config.Environment,
		},
	}
}
