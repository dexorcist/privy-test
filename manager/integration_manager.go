package manager

import (
	"privy-test/infra"
	"privy-test/integration/logging"
)

type integrationManager struct {
	infra infra.Infra
}

type IntegrationManager interface {
	Logger() logging.Logger
}

func NewIntegrationManager(infraInterface infra.Infra) IntegrationManager {
	return &integrationManager{infra: infraInterface}
}

func (i integrationManager) Logger() logging.Logger {
	return logging.NewLog(i.infra.Logger())
}
