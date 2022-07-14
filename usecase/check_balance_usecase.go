package usecase

import "livecode-lopei-grpc-client/repository"

type CheckBalanceUseCase interface {
	GetBalance(lopeiId int32) (float32, error)
}

type checkBalanceUseCase struct {
	repo repository.CustomerRepository
}

func (a *checkBalanceUseCase) GetBalance(lopeiId int32) (float32, error) {
	return a.repo.CheckBalance(lopeiId)
}

func NewCheckBalanceUseCase(repo repository.CustomerRepository) CheckBalanceUseCase {
	return &checkBalanceUseCase{
		repo: repo,
	}
}
