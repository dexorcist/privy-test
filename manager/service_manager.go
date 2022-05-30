package manager

import (
	"privy-test/infra"
	"privy-test/service"
)

type ServiceManager interface {
	HealthCheck() service.HealthCheckService
	Cake() service.CakeService
}

type serviceManager struct {
	infra       infra.Infra
	integration IntegrationManager
	repo        RepositoryManager
}

// NewServiceManager construct new service manager object which hold registered service singleton object
func NewServiceManager(infraInterface infra.Infra) ServiceManager {
	return &serviceManager{
		infra:       infraInterface,
		integration: NewIntegrationManager(infraInterface),
		repo:        NewRepositoryManager(infraInterface),
	}
}

func (s *serviceManager) HealthCheck() service.HealthCheckService {
	return service.NewHealthCheckService(s.infra.Config())
}

func (s *serviceManager) Cake() service.CakeService {
	return service.NewCakeService(s.repo.Cake(), s.integration.Logger())
}
