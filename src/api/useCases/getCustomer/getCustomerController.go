package getCustomer

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/api/useCases/getCustomer/presenter"
	"go-clean-architecture/src/application/useCases/GetCustomerUseCase"
)

type GetCustomerController struct {
	UseCase   GetCustomerUseCase.IGetCustomer
	Presenter *presenter.GetCustomerPresenter
}

func (controller *GetCustomerController) Get(c *gin.Context) {
	customerId := c.Param("customer_id")

	controller.UseCase.Get(customerId)

	c.IndentedJSON(controller.Presenter.Code, controller.Presenter.Response)
	return
}
