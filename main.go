package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(AnyErrorWrapper())
	r.GET("/", handleSome)
	log.Fatal(r.Run())
}

func handleSome(c *gin.Context) {
	c.Error(errors.New("What is this"))
	c.String(http.StatusOK, "Hello world!")
}

func AnyErrorWrapper() gin.HandlerFunc {
	return handleErrors(gin.ErrorTypeAny)
}

func handleErrors(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)
		if len(detectedErrors) > 0 {
			c.String(http.StatusInternalServerError, "An error occurred.")
			c.Abort()
			return
		}
	}
}
