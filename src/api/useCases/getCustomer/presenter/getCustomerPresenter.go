package presenter

import (
	"go-clean-architecture/src/api/useCases/getCustomer/io"
	"go-clean-architecture/src/domain"
	"net/http"
)

type IOutputPort interface {
	Success(obj domain.Customer)
	NotFound(message string)
	ValidationError(obj interface{}, err error)
}

type GetCustomerPresenter struct {
	Code     int
	Response interface{}
}

func (presenter *GetCustomerPresenter) Success(obj domain.Customer) {
	/* your http response logic goes here*/
	if obj.Id == "1" {

	}

	output := io.GetCustomerOutput{
		Id:   obj.Id,
		Name: obj.Name,
		Date: obj.BirthDate,
	}

	presenter.Response = output
	presenter.Code = http.StatusOK
}

func (presenter *GetCustomerPresenter) NotFound(message string) {
	/* your http response logic goes here*/
	presenter.Response = struct {
		Response string
	}{
		Response: message,
	}
	presenter.Code = http.StatusNotFound
}

func (presenter *GetCustomerPresenter) ValidationError(obj interface{}, err error) {
	/* your http response logic goes here*/
	presenter.Response = obj
	presenter.Code = http.StatusUnprocessableEntity
}
