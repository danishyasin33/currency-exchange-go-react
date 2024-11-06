package main

import (
	"currency-exchange/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/rates", routes.GetRates)

	// router.SetTrustedProxies("") // would set this for prod run to secure the API against unintended usage
	router.Run("localhost:8080")
}
