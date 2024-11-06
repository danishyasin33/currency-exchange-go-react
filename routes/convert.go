package routes

import (
	"log"
	"net/http"
	"reflect"
	"strconv"

	"currency-exchange/external"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type Error struct {
	Success bool
	Message string
}

func isValidCurrency(currency string) bool {
	switch currency {
	case
		"USD",
		"CAD",
		"MXN",
		"EUR",
		"GBP":
		return true
	}
	return false
}

func GetConvertHandler(c *cache.Cache) gin.HandlerFunc {
	coreFunc := func(context *gin.Context) {
		// there should be a better way to handle all three errors simultaneously so user can see all three at once instead of one by one. Not focusing on that for now

		currencyFrom := context.Query("from")
		if currencyFrom == "" {
			context.IndentedJSON(http.StatusBadRequest, &Error{Success: false, Message: "'from' currency not included in request"})
			return
		}

		if !isValidCurrency(currencyFrom) {
			context.IndentedJSON(http.StatusBadRequest, &Error{Success: false, Message: "'from' currency not a supported currency"})
			return
		}

		currencyTo := context.Query("to")
		if currencyTo == "" {
			context.IndentedJSON(http.StatusBadRequest, &Error{Success: false, Message: "'to' currency not included in request"})
			return
		}

		if !isValidCurrency(currencyTo) {
			context.IndentedJSON(http.StatusBadRequest, &Error{Success: false, Message: "'to' currency not a supported currency"})
			return
		}

		amount := context.Query("amount")
		if amount == "" {
			context.IndentedJSON(http.StatusBadRequest, &Error{Success: false, Message: "amount to be converted not included in request"})
			return
		}

		amountFloat, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			log.Fatal("something went wrong while converting amount to int")
			context.IndentedJSON(http.StatusInternalServerError, &Error{Success: false, Message: "amount cannot be converted to int"})
			return
		}

		cachedAllExchangeRates, found := c.Get("exchangeRates")
		if found {
			allExchangeRates := cachedAllExchangeRates.(*external.ConversionRates)

			// TODO: move this block to it's own method
			if currencyFrom == "USD" {
				reflection := reflect.ValueOf(allExchangeRates)
				currencyToValueByFieldName := reflect.Indirect(reflection).FieldByName(currencyTo)

				// TODO, make float more precise
				context.IndentedJSON(http.StatusOK, amountFloat*currencyToValueByFieldName.Float())
				return
			} else {
				reflection := reflect.ValueOf(allExchangeRates)
				currencyFromValueByFieldName := reflect.Indirect(reflection).FieldByName(currencyFrom)
				currencyToValueByFieldName := reflect.Indirect(reflection).FieldByName(currencyTo)

				// first convert from currency to USD
				fromCurrencyInUSD := amountFloat / currencyFromValueByFieldName.Float()

				context.IndentedJSON(http.StatusOK, fromCurrencyInUSD*currencyToValueByFieldName.Float())
				return
			}

		}

		updatedExchangeRate := external.GetExchangeRate()
		c.Set("exchangeRates", &updatedExchangeRate.ConversionRates, cache.DefaultExpiration)

		// TODO: move this block to it's own method, shares similar code to line 77
		if currencyFrom == "USD" {
			reflection := reflect.ValueOf(updatedExchangeRate.ConversionRates)
			valueByFieldName := reflect.Indirect(reflection).FieldByName(currencyTo)

			// TODO, make float more precise
			context.IndentedJSON(http.StatusOK, amountFloat*valueByFieldName.Float())
			return
		} else {
			reflection := reflect.ValueOf(updatedExchangeRate.ConversionRates)
			currencyFromValueByFieldName := reflect.Indirect(reflection).FieldByName(currencyFrom)
			currencyToValueByFieldName := reflect.Indirect(reflection).FieldByName(currencyTo)

			// first convert from currency to USD
			fromCurrencyInUSD := amountFloat / currencyFromValueByFieldName.Float()

			context.IndentedJSON(http.StatusOK, fromCurrencyInUSD*currencyToValueByFieldName.Float())
			return
		}
	}

	return gin.HandlerFunc(coreFunc)
}
