package manager

import "livecode-lopei-grpc-client/repository"

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
