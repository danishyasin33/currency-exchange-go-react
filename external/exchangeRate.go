package external

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ConversionRates struct {
	USD float32 `json:"USD"`
	CAD float32 `json:"CAD"`
	MXN float32 `json:"MXN"`
	EUR float32 `json:"EUR"`
	GBP float32 `json:"GBP"`
}

type ExchangeRateResponse struct {
	Result          string          `json:"result"`
	BaseCode        string          `json:"base_code"`
	ConversionRates ConversionRates `json:"conversion_rates"`
}

func GetExchangeRate() ExchangeRateResponse {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return ExchangeRateResponse{} // not a fan but just getting it done
	}

	exchangeRateApiKey := os.Getenv("EXCHANGE_RATE_API_KEY")
	requestUrl := "https://v6.exchangerate-api.com/v6/" + exchangeRateApiKey + "/latest/USD"

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return ExchangeRateResponse{} // not a fan but just getting it done
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return ExchangeRateResponse{} // not a fan but just getting it done
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return ExchangeRateResponse{} // not a fan but just getting it done
	}

	parsedResponse := &ExchangeRateResponse{}

	err = json.Unmarshal(resBody, &parsedResponse)

	if err != nil {
		fmt.Printf("client: could not parse body: %s\n", err)
		return ExchangeRateResponse{} // not a fan but just getting it done
	}

	return *parsedResponse
}
