package models

import "fmt"

type Premise struct {
	PremiseCode string
	Premise     string
	Address     string
	PremiseType string
	State       string
	District    string
}

func printPremiseList(premiseList []Premise) {
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

func CreatePremiseList(data [][]string) []Premise {
	var premiseList []Premise
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
	return premiseList
}
