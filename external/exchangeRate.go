package external

type ConversionRates struct {
	USD string `json:"USD"`
	CAD string `json:"CAD"`
	MXN string `json:"MXN"`
	EUR string `json:"EUR"`
	GBP string `json:"GBP"`
}

type ExchangeRateResponse struct {
	Result          string          `json:"result"`
	BaseCode        string          `json:"base_code"`
	ConversionRates ConversionRates `json:"conversion_rates"`
}

func getExchangeRate() {
}
