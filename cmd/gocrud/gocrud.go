package main

import (
	"accountservice/crud"
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"

	"github.com/pkg/errors"
	"github.com/unbxd/go-base/kit/transport/http"
	"github.com/unbxd/go-base/utils/log"
)

var gocrud *Gocrud

type (
	Option func(*Gocrud) error
	Gocrud struct {
		httpTransport *http.Transport
		logger        log.Logger
		crudControl   *crud.Control
	}
)

func (r *Gocrud) listen(transport *http.Transport, errch chan error) {
	err := transport.Open()
	if err != nil {
		errch <- errors.Wrap(err, "failed to start transport")
	}
}

func (r *Gocrud) Close(cx context.Context) error {
	r.logger.Flush()
	return nil
}

func (r *Gocrud) Open(cx context.Context) error {
	r.logger.Info("-- Starting Gocrud!")

	var (
		intchan = make(chan os.Signal, 1)
		errchan = make(chan error)
	)

	go r.listen(r.httpTransport, errchan)
	go signal.Notify(intchan, os.Interrupt)

	for {
		select {
		case <-intchan:
			r.logger.Info("Recieved os.Interrupt. Signal: ",
				log.String("signal", os.Interrupt.String()))
			r.logger.Info("Shutting down gracefully!!!")

			err := r.httpTransport.Close()
			if err != nil {
				panic(err)
			}

			r.logger.Info("Done!")
			return nil

		case err := <-errchan:
			r.logger.Error("Failed to start Gocrud server:", log.Error(err))
			return err
		}
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func NewGocrud(options ...Option) (*Gocrud, error) {
	var (
		tr, _ = http.NewTransport("0.0.0.0", "8000")
		lg, _ = log.NewZapLogger()
	)

	o := &Gocrud{
		httpTransport: tr,
		logger:        lg,
	}

	for _, ofn := range options {
		fmt.Println(">> ----- Initializing: ", getFunctionName(ofn))
		err := ofn(o)
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}
