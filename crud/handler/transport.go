package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"accountservice/crud/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

func MakeCreateCustomerEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequest)
		msg, err := s.CreateCustomer(ctx, req.customer)
		return CreateCustomerResponse{Msg: msg, Err: err}, nil
	}
}

func MakeGetCustomerByIdEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByIdRequest)
		customerDetails, err := s.GetCustomerById(ctx, req.Id)
		if err != nil {
			return GetCustomerByIdResponse{Customer: customerDetails, Err: "Id not found"}, nil
		}
		return GetCustomerByIdResponse{Customer: customerDetails, Err: ""}, nil
	}
}

func MakeGetAllCustomersEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		customerDetails, err := s.GetAllCustomers(ctx)
		if err != nil {
			return GetAllCustomersResponse{Customer: customerDetails, Err: "no data found"}, nil
		}
		return GetAllCustomersResponse{Customer: customerDetails, Err: ""}, nil
	}
}

func MakeDeleteCustomerEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCustomerRequest)
		msg, err := s.DeleteCustomer(ctx, req.Customerid)
		if err != nil {
			return DeleteCustomerResponse{Msg: msg, Err: err}, nil
		}
		return DeleteCustomerResponse{Msg: msg, Err: nil}, nil
	}
}
func MakeUpdateCustomerendpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCustomerRequest)
		msg, err := s.UpdateCustomer(ctx, req.customer)
		return msg, err
	}
}

func DecodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateCustomerRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetCustomerByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetCustomerByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetCustomerByIdRequest{
		Id: vars["customerid"],
	}
	return req, nil
}
func DecodeGetAllCustomersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into GETALL Decoding")
	var req GetAllCustomersRequest
	return req, nil
}
func DecodeDeleteCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteCustomerRequest
	vars := mux.Vars(r)
	req = DeleteCustomerRequest{
		Customerid: vars["customerid"],
	}
	return req, nil
}
func DecodeUpdateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

type (
	CreateCustomerRequest struct {
		customer service.Customer
	}
	CreateCustomerResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	GetCustomerByIdRequest struct {
		Id string `json:"customerid"`
	}
	GetCustomerByIdResponse struct {
		Customer interface{} `json:"customer,omitempty"`
		Err      string      `json:"error,omitempty"`
	}
	GetAllCustomersRequest struct{}

	GetAllCustomersResponse struct {
		Customer interface{} `json:"customer,omitempty"`
		Err      string      `json:"error,omitempty"`
	}
	DeleteCustomerRequest struct {
		Customerid string `json:"customerid"`
	}

	DeleteCustomerResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
	UpdateCustomerRequest struct {
		customer service.Customer
	}
	UpdateCustomerResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
)
