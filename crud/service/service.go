package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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

// Describes the AccountServiceStruct for repository interaction
type Repository interface {
	CreateCustomer(ctx context.Context, customer Customer) error
	GetCustomerById(ctx context.Context, id string) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)
	UpdateCustomer(ctx context.Context, customer Customer) (string, error)
	DeleteCustomer(ctx context.Context, id string) (string, error)
}

type AccountServiceStruct struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) AccountService {
	return &AccountServiceStruct{
		repository: rep,
		logger:     logger,
	}
}

func (s AccountServiceStruct) CreateCustomer(ctx context.Context, customer Customer) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	var msg = "success"
	customerDetails := Customer{
		Customerid: customer.Customerid,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}
	if err := s.repository.CreateCustomer(ctx, customerDetails); err != nil {
		level.Error(logger).Log("err from repo is ", err)
		return "", err
	}
	return msg, nil
}

func (s AccountServiceStruct) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetcustomerById")

	var customer interface{}
	var empty interface{}
	customer, err := s.repository.GetCustomerById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return customer, nil
}

func (s AccountServiceStruct) GetAllCustomers(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetAllcustomers")
	var customer interface{}
	var empty interface{}
	customer, err := s.repository.GetAllCustomers(ctx)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return customer, nil
}

func (s AccountServiceStruct) DeleteCustomer(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteCustomer")
	msg, err := s.repository.DeleteCustomer(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return "", err
	}
	return msg, nil
}

func (s AccountServiceStruct) UpdateCustomer(ctx context.Context, customer Customer) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	var msg = "success"
	customerDetails := Customer{
		Customerid: customer.Customerid,
		Email:      customer.Email,
		Phone:      customer.Phone,
	}
	msg, err := s.repository.UpdateCustomer(ctx, customerDetails)
	if err != nil {
		level.Error(logger).Log("err from repo is ", err)
		return "", err
	}
	return msg, nil
}
