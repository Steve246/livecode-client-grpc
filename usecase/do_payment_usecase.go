package usecase

import "livecode-lopei-grpc-client/repository"

type DoPaymentUseCase interface {
	DoPayment(lopeiId int32, amount float32) error
}

type doPaymentUseCase struct {
	repo repository.CustomerRepository
}