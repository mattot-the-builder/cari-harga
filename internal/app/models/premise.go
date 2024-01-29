package models

import (
	"strings"

	"github.com/mattot-the-builder/go-csv/internal/app/utils"
)

type Premise struct {
	PremiseCode string
	Premise     string
	Address     string
	PremiseType string
	State       string
	District    string
}

var premiseList []Premise

func init() {
	CreatePremiseList(utils.ReadCsvFile("data/lookup_premise.csv"))
}

func GetPremiseList() []Premise {
	return premiseList
}

func SearchPremise(keyword string) []Premise {
	var searchResult []Premise
	for _, premise := range premiseList {
		if strings.Contains(premise.Premise, strings.ToUpper(keyword)) {
			searchResult = append(searchResult, premise)
		}
	}
	return searchResult
}

func CreatePremiseList(data [][]string) { 
	for i, line := range data {
		if i > 0 {
			var rec Premise

			for j, field := range line {
				switch j {
				case 0:
					rec.PremiseCode = field
					break
				case 1:
					rec.Premise = field
					break
				case 2:
					rec.Address = field
					break
				case 3:
					rec.PremiseType = field
					break
				case 4:
					rec.State = field
					break
				case 5:
					rec.District = field
					break
				default:
					break
				}
			}

			premiseList = append(premiseList, rec)
		}

	}
}
