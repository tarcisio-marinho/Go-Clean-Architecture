package middleware

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/application/shared"
	"log"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header["Authorization"]

		if len(authorization) != 1 {
			c.AbortWithStatusJSON(http.StatusBadRequest, "without authorization header")
			return
		}

		jwt := authorization[0]

		if shared.IsNullOrEmpty(jwt) {
			c.AbortWithStatusJSON(http.StatusBadRequest, "empty authorization header")
			return
		}

		valid := ValidateJwt(jwt)
		if !valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid jwt")
			return
		}

		c.Next()
	}
}

func ValidateJwt(jwt string) bool {
	/* Validate jwt logic goes here */
	return true
}
