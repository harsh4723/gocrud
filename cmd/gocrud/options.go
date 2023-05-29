package main

import (
	"accountservice/crud"
	"accountservice/utils"

	"github.com/pkg/errors"
	"github.com/unbxd/go-base/kit/transport/http"
	"github.com/unbxd/go-base/utils/log"
)

func WithCustomLogger(
	level string,
	encoding string,
	output string,
) Option {
	return func(r *Gocrud) (err error) {
		logger, err := log.NewZapLogger(
			log.ZapWithLevel(level),
			log.ZapWithEncoding(encoding),
			log.ZapWithOutput([]string{output}),
		)

		if err != nil {
			return errors.Wrap(err, "failed to create logger")
		}

		r.logger = logger
		return
	}
}

func WithHTTPTransport(
	host, port string,
	monitor []string,
) Option {
	return func(r *Gocrud) (err error) {
		tr, err := http.NewTransport(
			host,
			port,
			http.WithLogger(r.logger),
			http.WithFullDefaults(),
			http.WithMonitors(monitor),
		)
		if err != nil {
			return err
		}

		r.httpTransport = tr
		return
	}
}

func WithCrudControlPlane() Option {
	return func(r *Gocrud) (err error) {
		db := utils.GetDBconn()
		repository, err := crud.NewRepo(db, r.logger)
		rc := crud.NewControl(r.logger, repository)

		r.crudControl = rc

		r.crudControl.Bind(
			r.httpTransport,
		)

		return
	}
}
