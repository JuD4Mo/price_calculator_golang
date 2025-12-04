package prices

import (
	"fmt"

	"example.com/price-calculator/IIOManager"
	"example.com/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	IOManager         IIOManager.IOManager `json:"-"`
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

func (job TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.readPrices()
	if err != nil {
		// return err
		errorChan <- err
		return
	}
	result := make(map[string]string)
	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom IIOManager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		Prices:    []float64{10, 20, 30},
		TaxRate:   taxRate,
	}
}
