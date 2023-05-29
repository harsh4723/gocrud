package main

import (
	"accountservice/cmd/flags"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	err := (&cli.App{
		Name:     "gocrud",
		Usage:    "Simple crud apis",
		Version:  "1.0",
		Before:   badge,
		Flags:    flags.Flags(),
		Commands: commands,
	}).Run(os.Args)
	if err != nil {
		fmt.Println("Something Went Wrong. Failed to start Gocrud.: " + err.Error())
		log.Fatal(
			fmt.Sprintf(
				"-- \nfailed to start Gocrud. \n--\n Caused By:\n%s\n--",
				errorstack(err.Error()),
			),
		)
	}
	// 	logger := log.NewLogfmtLogger(os.Stderr)
	// 	db := utils.GetDBconn()

	// 	r := mux.NewRouter()

	// 	var svc service.AccountService
	// 	svc = service.AccountServiceStruct{}
	// 	{
	// 		repository, err := repo.NewRepo(db, logger)
	// 		if err != nil {
	// 			level.Error(logger).Log("exit", err)
	// 			os.Exit(-1)
	// 		}
	// 		svc = service.NewService(repository, logger)
	// 	}

	// 	CreateAccountHandler := httptransport.NewServer(
	// 		handler.MakeCreateCustomerEndpoint(svc),
	// 		handler.DecodeCreateCustomerRequest,
	// 		handler.EncodeResponse,
	// 	)

	// 	GetByCustomerIdHandler := httptransport.NewServer(
	// 		handler.MakeGetCustomerByIdEndpoint(svc),
	// 		handler.DecodeGetCustomerByIdRequest,
	// 		handler.EncodeResponse,
	// 	)
	// 	GetAllCustomersHandler := httptransport.NewServer(
	// 		handler.MakeGetAllCustomersEndpoint(svc),
	// 		handler.DecodeGetAllCustomersRequest,
	// 		handler.EncodeResponse,
	// 	)
	// 	DeleteCustomerHandler := httptransport.NewServer(
	// 		handler.MakeDeleteCustomerEndpoint(svc),
	// 		handler.DecodeDeleteCustomerRequest,
	// 		handler.EncodeResponse,
	// 	)
	// 	UpdateCustomerHandler := httptransport.NewServer(
	// 		handler.MakeUpdateCustomerendpoint(svc),
	// 		handler.DecodeUpdateCustomerRequest,
	// 		handler.EncodeResponse,
	// 	)

	// 	http.Handle("/", r)
	// 	http.Handle("/account", CreateAccountHandler)
	// 	http.Handle("/account/update", UpdateCustomerHandler)
	// 	r.Handle("/account/getAll", GetAllCustomersHandler).Methods("GET")
	// 	r.Handle("/account/{customerid}", GetByCustomerIdHandler).Methods("GET")
	// 	r.Handle("/account/{customerid}", DeleteCustomerHandler).Methods("DELETE")

	// // http.Handle("/metrics", promhttp.Handler())
	// logger.Log("msg", "HTTP", "addr", ":8000")
	// logger.Log("err", http.ListenAndServe(":8000", nil))
}
