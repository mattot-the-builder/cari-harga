package controllers

import (
	"fmt"

	"github.com/mattot-the-builder/go-csv/internal/app/models"
)


func PrintPremiseList(premiseList []models.Premise) {
	for _, premise := range premiseList {
		fmt.Println("Premise data start")
		fmt.Println(premise.PremiseCode)
		fmt.Println(premise.Premise)
		fmt.Println(premise.Address)
		fmt.Println(premise.PremiseType)
		fmt.Println(premise.State)
		fmt.Println(premise.District)
		fmt.Printf("Premise data end\n\n")
	}
}

