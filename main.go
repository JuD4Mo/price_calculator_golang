package main

import (
	"fmt"

	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, tax := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)

		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tax*100))
		//cmd := cmdmanager.New()
		taxJob := prices.NewTaxIncludedPriceJob(fm, tax)
		go taxJob.Process(doneChans[i], errorChans[i])

		// if err != nil {
		// 	fmt.Println("could not process job")
		// 	fmt.Println(err)
		// }
	}

	/*
		Select sirve para manejar casos en donde las go rutines pueden pasar diferentes valores
		En este caso, son manejados, tanto los errores como mensajes de éxito.
	*/
	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}
	}
}
