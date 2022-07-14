package manager

import (
	"livecode-lopei-grpc-client/usecase"
)

type UseCaseManager interface {
	CheckBalanceUseCase() usecase.CheckBalanceUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (s *useCaseManager) CheckBalanceUseCase() usecase.CheckBalanceUseCase {
	return usecase.NewCheckBalanceUseCase(s.repoManager.CustomerRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
