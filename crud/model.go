package crud

type (
	CreateCustomerRequest struct {
		customer Customer
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
		customer Customer
	}
	UpdateCustomerResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
)
