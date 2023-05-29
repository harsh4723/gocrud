package main

import (
	"net/http"
	"os"

	"accountservice/crud/service"
	"accountservice/utils"

	"accountservice/crud/repo"

	"accountservice/crud/handler"

	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db := utils.GetDBconn()

	r := mux.NewRouter()

	var svc service.AccountService
	svc = service.AccountServiceStruct{}
	{
		repository, err := repo.NewRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = service.NewService(repository, logger)
	}

	CreateAccountHandler := httptransport.NewServer(
		handler.MakeCreateCustomerEndpoint(svc),
		handler.DecodeCreateCustomerRequest,
		handler.EncodeResponse,
	)

	GetByCustomerIdHandler := httptransport.NewServer(
		handler.MakeGetCustomerByIdEndpoint(svc),
		handler.DecodeGetCustomerByIdRequest,
		handler.EncodeResponse,
	)
	GetAllCustomersHandler := httptransport.NewServer(
		handler.MakeGetAllCustomersEndpoint(svc),
		handler.DecodeGetAllCustomersRequest,
		handler.EncodeResponse,
	)
	DeleteCustomerHandler := httptransport.NewServer(
		handler.MakeDeleteCustomerEndpoint(svc),
		handler.DecodeDeleteCustomerRequest,
		handler.EncodeResponse,
	)
	UpdateCustomerHandler := httptransport.NewServer(
		handler.MakeUpdateCustomerendpoint(svc),
		handler.DecodeUpdateCustomerRequest,
		handler.EncodeResponse,
	)

	http.Handle("/", r)
	http.Handle("/account", CreateAccountHandler)
	http.Handle("/account/update", UpdateCustomerHandler)
	r.Handle("/account/getAll", GetAllCustomersHandler).Methods("GET")
	r.Handle("/account/{customerid}", GetByCustomerIdHandler).Methods("GET")
	r.Handle("/account/{customerid}", DeleteCustomerHandler).Methods("DELETE")

	// http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
