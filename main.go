package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	initialTime := time.Now()
	f, err := os.Open("lookup_item.csv")
	if err != nil {
		log.Println("Error opening file ", err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Println("Error reading csv file ", err)
	}

	// premiseList := createPremiseList(data)
	itemList := createItemList(data)

	for _, item := range itemList {
		fmt.Println("Premise data start")
		fmt.Println(item.ItemName)
		fmt.Println(item.ItemCategory)
		fmt.Println(item.UnitOfMeasurement)
		fmt.Printf("Premise data end\n\n")
	}

	fmt.Printf("Total rows iterated: %v\n", len(itemList))
	fmt.Printf("Time taken: %v\n", time.Now().Sub(initialTime))
}
