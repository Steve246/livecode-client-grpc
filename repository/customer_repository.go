package repository

import (
	"context"
	"enigmacamp.com/lopei_grpc_cnlt/service"
	"fmt"
	"log"
)

type CustomerRepository interface {
	CheckBalance(lopeiId int32) (float32, error)
	DoPayment(lopeiId int32, amount float32) error
}

type customerRepository struct {
	client service.LopeiPaymentClient
}

func (c *customerRepository) CheckBalance(lopeiId int32) (float32, error) {
	response, err := c.client.CheckBalance(context.Background(), &service.CheckBalanceMessage{
		LopeiId: lopeiId,
	})
	if err != nil {
		return 0, err
	}
	fmt.Println(response)
	return 0, nil
}

func (c *customerRepository) DoPayment(lopeiId int32, amount float32) error {
	response, err := c.client.DoPayment(context.Background(), &service.PaymentMessage{
		LopeiId: lopeiId,
		Amount:  amount,
	})
	if err != nil {
		return err
	}
	log.Printf("%v", response)
	return nil
}

func NewCustomerRepository(client service.LopeiPaymentClient) CustomerRepository {
	repo := new(customerRepository)
	repo.client = client
	return repo
}
