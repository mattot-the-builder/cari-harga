package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mattot-the-builder/go-csv/internal/app/models"
)

func GetAllPrices(w http.ResponseWriter, r *http.Request) {
	priceCatcherList := models.GetPriceCatcherList()
	initialTime := time.Now()
	for i, priceCatcher := range priceCatcherList {
		log.Printf("Data number: %v\n", i)
		fmt.Println("Price data start")
		fmt.Printf("Item code: %v\n", priceCatcher.ItemCode)
		fmt.Printf("Premise code: %v\n", priceCatcher.PremiseCode)
		fmt.Printf("Item price: %v\n", priceCatcher.Price)
		fmt.Printf("Item data end\n\n")
	}

	durationTaken := time.Now().Sub(initialTime)
	fmt.Fprintf(w, "Total rows iterated: %v\n", len(priceCatcherList))
	log.Printf("Total rows iterated: %v\n", len(priceCatcherList))
	fmt.Fprintf(w, "Time taken: %v\n", durationTaken)
	log.Printf("Time taken: %v\n", durationTaken)

	averageReadSpeed := float64(len(priceCatcherList)) / float64(durationTaken.Abs().Seconds())
	log.Printf("Average data read per second: %.2f\n rows/s", averageReadSpeed)
}
