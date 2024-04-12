package taxes

type TaxData struct {
	TaxRates map[string]float64 `json:"tax_rate"`
}
