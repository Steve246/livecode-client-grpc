package manager

import "enigmacamp.com/lopei_grpc_cnlt/repository"

type RepositoryManager interface {
	CustomerRepo() repository.CustomerRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infraManager.LopeiClientConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
