package manager

import (
	"privy-test/infra"
	"privy-test/repository"
	"privy-test/utils"
)

type RepositoryManager interface {
	Cake() repository.CakeRepository
}

type repositoryManager struct {
	infra infra.Infra
}

func (r repositoryManager) Cake() repository.CakeRepository {
	return repository.NewCakeRepository(r.infra.SQLDB(), utils.NewCommonService())
}

func NewRepositoryManager(infraInterface infra.Infra) RepositoryManager {
	return repositoryManager{
		infra: infraInterface,
	}
}
