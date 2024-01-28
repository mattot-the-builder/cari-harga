package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mattot-the-builder/go-csv/internal/app/controllers"
	"github.com/mattot-the-builder/go-csv/internal/app/models"
)

func readCsvFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening file ", err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Println("Error reading csv file ", err)
	}
	return data
}

func main() {
	initialTime := time.Now()

	data := readCsvFile("data/lookup_premise.csv")

	premiseList := models.CreatePremiseList(data)

	controllers.PrintPremiseList(premiseList)

	fmt.Printf("Total rows iterated: %v\n", len(premiseList))
	fmt.Printf("Time taken: %v\n", time.Now().Sub(initialTime))
}
