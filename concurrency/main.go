package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([][]chan any, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = []chan any{make(chan any), make(chan any)}
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index][0], doneChans[index][1])
	}

	for index, chans := range doneChans {
		select {
		case err := <-chans[1]:
			if err != nil {
				fmt.Printf("Job %d failed: %v\n", index, err)
			}
		case <-chans[0]:
			fmt.Printf("Job %d completed successfully\n", index)
		}
	}
	fmt.Println("All jobs completed")
}
