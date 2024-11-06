package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type Error struct {
	Success bool
	Message string
}

func GetConvertHandler(c *cache.Cache) gin.HandlerFunc {
	coreFunc := func(context *gin.Context) {
		// there should be a better way to handle all three errors simultaneously so user can see all three at once instead of one by one. Not focusing on that for now

		currencyFrom := context.Query("from")
		if currencyFrom == "" {
			context.IndentedJSON(http.StatusBadRequest, &Error{Success: false, Message: "'from' currency not included in request"})
			return
		}

		fmt.Println(currencyFrom)

		currencyTo := context.Query("to")
		if currencyTo == "" {
			context.IndentedJSON(http.StatusBadRequest, &Error{Success: false, Message: "'to' currency not included in request"})
			return
		}

		amount := context.Query("amount")
		if amount == "" {
			context.IndentedJSON(http.StatusBadRequest, &Error{Success: false, Message: "amount to be converted not included in request"})
			return
		}

		// exchangeRates, found := c.Get("exchangeRates")
		// if found {
		// 	context.IndentedJSON(http.StatusOK, exchangeRates)
		// 	return
		// }

		// updatedExchangeRate := external.GetExchangeRate()

		// c.Set("exchangeRates", &updatedExchangeRate.ConversionRates, cache.DefaultExpiration)

	}

	return gin.HandlerFunc(coreFunc)
}
