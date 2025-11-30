package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iiomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iiomanager.IOManager `json:"-"`
	TaxRate           float64              `json:"tax_rate"`
	Prices            []float64            `json:"input_prices"`
	TaxIncludedPrices map[string]string    `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) readPrices() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}
	job.Prices = prices

	return nil
}

func (job TaxIncludedPriceJob) Process() error {
	err := job.readPrices()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iiomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		Prices:    []float64{10, 20, 30},
		TaxRate:   taxRate,
	}
}
