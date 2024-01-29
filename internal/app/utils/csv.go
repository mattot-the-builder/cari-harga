package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(fileName string) [][]string {
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

