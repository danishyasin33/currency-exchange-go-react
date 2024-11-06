package main

import (
	"currency-exchange/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/patrickmn/go-cache"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Create a cache with a default expiration time of 10 minutes, and which
	// purges expired items every 10 minutes
	cache := cache.New(10*time.Minute, 10*time.Minute)

	router := gin.Default()
	router.GET("/rates", routes.GetRatesHandler(cache))
	router.GET("/convert", routes.GetConvertHandler(cache))

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := setupRouter()

	router.Use(cors.Default())

	// router.SetTrustedProxies("") // would set this for prod run to secure the API against unintended usage
	router.Run(":8080")
}
