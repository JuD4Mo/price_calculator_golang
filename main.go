package main

import (
	"fmt"

	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxrates := []float64{0, 0.7, 0.1, 0.15}

	for _, tax := range taxrates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		//cmd := cmdmanager.New()
		taxJob := prices.NewTaxIncludedPriceJob(fm, tax)
		err := taxJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}
}
