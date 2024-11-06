package main

import (
	"currency-exchange/routes"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a cache with a default expiration time of 10 minutes, and which
	// purges expired items every 10 minutes
	cache := cache.New(10*time.Minute, 10*time.Minute)

	router := gin.Default()
	router.GET("/rates", routes.GetRatesHandler(cache))
	router.GET("/convert", routes.GetConvertHandler(cache))

	// router.SetTrustedProxies("") // would set this for prod run to secure the API against unintended usage
	router.Run("localhost:8080")
}
