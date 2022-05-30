package service

import (
	"context"
	"privy-test/param/cake"
	"privy-test/param/healthcheck"
)

type HealthCheckService interface {
	HealthCheck(ctx context.Context) healthcheck.HTTPHealthCheckResponse
}

type CakeService interface {
	Create(ctx context.Context, request *cake.CreateUpdateRequest) (*cake.HTTPGetDetailCakeResponse, error)
	Update(ctx context.Context, cakeID int64, request *cake.CreateUpdateRequest) (*cake.HTTPGetDetailCakeResponse, error)
	Delete(ctx context.Context, cakeID int64) error
	GetDetail(ctx context.Context, cakeID int64) (*cake.HTTPGetDetailCakeResponse, error)
	GetList(ctx context.Context, request *cake.FindAllRequest) (*cake.HTTPGetListCakeResponse, error)
}
