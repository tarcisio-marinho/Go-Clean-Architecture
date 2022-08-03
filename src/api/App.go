package api

import (
	"go-clean-architecture/src/api/configs"
	"go-clean-architecture/src/api/useCases/getCustomer"
	"go-clean-architecture/src/api/useCases/getCustomer/presenter"
	"go-clean-architecture/src/application/useCases/GetCustomerUseCase"
	"go-clean-architecture/src/infrastructure/repository"
)

type App struct {
	/* Add here your controllers */
	Controller getCustomer.GetCustomerController
}

func BuildApp() (App, error) { /* You can change to dependency injection here */
	db := repository.GetCustomerQuery{
		Database: repository.CreateDatabase{
			DbConfig:   configs.GetDatabaseConfig(),
			Database:   "database",
			Collection: "collection",
		},
	}

	presenter := presenter.GetCustomerPresenter{}
	controller := getCustomer.GetCustomerController{
		UseCase: GetCustomerUseCase.GetCustomerUseCase{
			OutputPort: &presenter,
			Query:      db,
		},
		Presenter: &presenter,
	}

	return App{
		Controller: controller,
	}, nil
}
