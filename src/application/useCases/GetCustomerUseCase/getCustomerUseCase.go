package GetCustomerUseCase

import (
	"errors"
	"fmt"
	"go-clean-architecture/src/api/useCases/getCustomer/presenter"
	"go-clean-architecture/src/application/shared"
	"go-clean-architecture/src/infrastructure/repository"
)

type IGetCustomer interface {
	Get(string)
}

type GetCustomerUseCase struct {
	OutputPort presenter.IOutputPort
	Query      repository.IGetCustomerQuery
}

func (useCase GetCustomerUseCase) Get(customerId string) {

	if shared.IsNullOrEmpty(customerId) {
		fmt.Println("orderId is empty")
		useCase.OutputPort.ValidationError("error validating, customerId is empty", errors.New("empty customer id"))
		return
	}

	output, err := useCase.Query.Get(customerId)

	if err != nil {
		fmt.Println("Order not found:", err)
		useCase.OutputPort.NotFound("Order not found")
		return
	}

	useCase.OutputPort.Success(output)
}
