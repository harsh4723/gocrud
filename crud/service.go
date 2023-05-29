package crud

import (
	"context"

	"github.com/unbxd/go-base/utils/log"
)

type Customer struct {
	Customerid string `json:"customerid"`
	Email      string ` json:"email"`
	Phone      string ` json:"phone"`
}

type AccountService interface {
	CreateCustomer(ctx context.Context, customer Customer) (string, error)
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)
	UpdateCustomer(ctx context.Context, customer Customer) (string, error)
	DeleteCustomer(ctx context.Context, id string) (string, error)
}

type Repository interface {
	CreateCustomer(ctx context.Context, customer Customer) error
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)
	UpdateCustomer(ctx context.Context, customer Customer) (string, error)
	DeleteCustomer(ctx context.Context, id string) (string, error)
}

type AccountServiceStruct struct {
	logger     log.Logger
	repository Repository
}

func (s AccountServiceStruct) CreateCustomer(ctx context.Context, customer Customer) (string, error) {
	logger := s.logger
	var msg = "success"
	customerDetails := Customer{
		Customerid: customer.Customerid,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}
	if err := s.repository.CreateCustomer(ctx, customerDetails); err != nil {
		logger.Error("create customer err from repo is ", log.Error(err))
		return "", err
	}
	return msg, nil
}

func (s AccountServiceStruct) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	logger := s.logger

	var customer interface{}
	var empty interface{}
	customer, err := s.repository.GetCustomerById(ctx, id)
	if err != nil {
		logger.Error("get customer by id err from repo is ", log.Error(err))
		return empty, err
	}
	return customer, nil
}

func (s AccountServiceStruct) GetAllCustomers(ctx context.Context) (interface{}, error) {
	logger := s.logger
	var customer interface{}
	var empty interface{}
	customer, err := s.repository.GetAllCustomers(ctx)
	if err != nil {
		logger.Error("get all customer err from repo is ", log.Error(err))
		return empty, err
	}
	return customer, nil
}

func (s AccountServiceStruct) DeleteCustomer(ctx context.Context, id string) (string, error) {
	logger := s.logger
	msg, err := s.repository.DeleteCustomer(ctx, id)
	if err != nil {
		logger.Error("delete customer by id err from repo is ", log.Error(err))
		return "", err
	}
	return msg, nil
}

func (s AccountServiceStruct) UpdateCustomer(ctx context.Context, customer Customer) (string, error) {
	logger := s.logger
	var msg = "success"
	customerDetails := Customer{
		Customerid: customer.Customerid,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}
	msg, err := s.repository.UpdateCustomer(ctx, customerDetails)
	if err != nil {
		logger.Error("update customer err from repo is ", log.Error(err))
		return "", err
	}
	return msg, nil
}

func newSvc(
	l log.Logger,
	r Repository,
) AccountService {

	return &AccountServiceStruct{l, r}
}
