package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	middleware "go-clean-architecture/src/api/midleware"
)

func Server() {
	router := gin.Default()

	router.Use(middleware.CheckJWT())
	router.Use(middleware.Logger())

	app, err := BuildApp()
	if err != nil {
		panic("error building the app")
	}

	router.GET("/v2/customer/:customer_id", app.Controller.Get)

	addr := fmt.Sprintf("0.0.0.0:%v", 3001)
	router.Run(addr)
}
