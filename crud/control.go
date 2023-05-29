package crud

import (
	"github.com/unbxd/go-base/kit/transport/http"
	"github.com/unbxd/go-base/utils/log"
)

type Control struct {
	service AccountService
}

func (c *Control) Service() AccountService { return c.service }

func (c *Control) Bind(ht *http.Transport, opts ...http.HandlerOption) {
	ht.POST(
		"/account",
		CreateAccountHandler(c.service),
		MakeCreateAccountHanlerOption(opts)...,
	)

	ht.GET(
		"/account/getAll",
		GetAllCustomersHandler(c.service),
		MakeGetAllCustomersHandlerOption(opts)...,
	)
}

func NewControl(
	l log.Logger,
	r Repository,
) *Control {
	svc := newSvc(l, r)
	return &Control{svc}
}
