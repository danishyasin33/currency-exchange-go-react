package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type CurrencyType struct {
	ID      string  `json:"id"` // ids would generally be UUIDs but keeping them as string for now
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	UsdRate float64 `json:"usdrate"` // could probably use float32 here
}

// specified currencies that we'll be supporting
// defaulting USD rate to 0
var currencies = []CurrencyType{
	{ID: "1", Name: "Canadian Dollar", Symbol: "CAD", UsdRate: 0},
	{ID: "2", Name: "Mexican Peso", Symbol: "MXN", UsdRate: 0},
	{ID: "3", Name: "Euro", Symbol: "EUR", UsdRate: 0},
	{ID: "4", Name: "British Pound", Symbol: "GBP", UsdRate: 0},
}

// gets the latest currency exchange rates if cache is invalid
func GetRates(context *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	exchangeRateApiKey := os.Getenv("EXCHANGE_RATE_API_KEY")

	println(exchangeRateApiKey)

	context.IndentedJSON(http.StatusOK, currencies)
}
