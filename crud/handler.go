package crud

import (
	"context"
	"encoding/json"
	"fmt"

	net_http "net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/unbxd/go-base/kit/transport/http"
)

func MakeCreateCustomerEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequest)
		msg, err := s.CreateCustomer(ctx, req.customer)
		return CreateCustomerResponse{Msg: msg, Err: err}, nil
	}
}

func MakeGetCustomerByIdEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByIdRequest)
		customerDetails, err := s.GetCustomerById(ctx, req.Id)
		if err != nil {
			return GetCustomerByIdResponse{Customer: customerDetails, Err: "Id not found"}, nil
		}
		return GetCustomerByIdResponse{Customer: customerDetails, Err: ""}, nil
	}
}

func MakeGetAllCustomersEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		customerDetails, err := s.GetAllCustomers(ctx)
		if err != nil {
			return GetAllCustomersResponse{Customer: customerDetails, Err: "no data found"}, nil
		}
		return GetAllCustomersResponse{Customer: customerDetails, Err: ""}, nil
	}
}

func MakeDeleteCustomerEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCustomerRequest)
		msg, err := s.DeleteCustomer(ctx, req.Customerid)
		if err != nil {
			return DeleteCustomerResponse{Msg: msg, Err: err}, nil
		}
		return DeleteCustomerResponse{Msg: msg, Err: nil}, nil
	}
}
func MakeUpdateCustomerendpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCustomerRequest)
		msg, err := s.UpdateCustomer(ctx, req.customer)
		return msg, err
	}
}

func CreateAccountHandler(service AccountService) http.Handler {
	return http.Handler(MakeCreateCustomerEndpoint(service))
}

func GetByCustomerIdHandler(service AccountService) http.Handler {
	return http.Handler(MakeGetCustomerByIdEndpoint(service))
}

func GetAllCustomersHandler(service AccountService) http.Handler {
	return http.Handler(MakeGetAllCustomersEndpoint(service))
}

func DeleteCustomerHandler(service AccountService) http.Handler {
	return http.Handler(MakeDeleteCustomerEndpoint(service))
}

func UpdateCustomerHandler(service AccountService) http.Handler {
	return http.Handler(MakeUpdateCustomerendpoint(service))
}

func MakeCreateAccountHanlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(DecodeCreateCustomerRequest),
		http.HandlerWithEncoder(EncodeResponse),
	}, opts...)
}

func MakeGetAllCustomersHandlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(DecodeGetAllCustomersRequest),
		http.HandlerWithEncoder(EncodeResponse),
	}, opts...)
}

func MakeGetCustomerByIdHandlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(DecodeGetCustomerByIdRequest),
		http.HandlerWithEncoder(EncodeResponse),
	})
}

func MakeDeleteCustomerHandlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(DecodeDeleteCustomerRequest),
		http.HandlerWithEncoder(EncodeResponse),
	})
}

func MakeUpdateCustomerHandlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(DecodeUpdateCustomerRequest),
		http.HandlerWithEncoder(EncodeResponse),
	})
}

func DecodeCreateCustomerRequest(_ context.Context, r *net_http.Request) (interface{}, error) {
	var req CreateCustomerRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetCustomerByIdRequest(_ context.Context, r *net_http.Request) (interface{}, error) {
	var req GetCustomerByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	// vars := mux.Vars(r)
	customerid := http.Parameters(r).ByName("customerid")
	req = GetCustomerByIdRequest{
		Id: customerid,
	}
	return req, nil
}
func DecodeGetAllCustomersRequest(_ context.Context, r *net_http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into GETALL Decoding")
	var req GetAllCustomersRequest
	return req, nil
}
func DecodeDeleteCustomerRequest(_ context.Context, r *net_http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteCustomerRequest
	//vars := mux.Vars(r)
	customerid := http.Parameters(r).ByName("customerid")
	req = DeleteCustomerRequest{
		Customerid: customerid,
	}
	return req, nil
}
func DecodeUpdateCustomerRequest(_ context.Context, r *net_http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w net_http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}
