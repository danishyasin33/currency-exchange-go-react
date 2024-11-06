package routes

import (
	"net/http"

	"currency-exchange/external"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// gets the latest currency exchange rates if cache is invalid
func GetRatesHandler(c *cache.Cache) gin.HandlerFunc {
	coreFunc := func(context *gin.Context) {
		exchangeRates, found := c.Get("exchangeRates")
		if found {
			context.IndentedJSON(http.StatusOK, exchangeRates)
			return
		}

		updatedExchangeRate := external.GetExchangeRate()

		c.Set("exchangeRates", &updatedExchangeRate.ConversionRates, cache.DefaultExpiration)
		context.IndentedJSON(http.StatusOK, updatedExchangeRate.ConversionRates)
	}

	return gin.HandlerFunc(coreFunc)
}
