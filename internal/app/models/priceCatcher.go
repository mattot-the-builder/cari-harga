package models

import "github.com/mattot-the-builder/go-csv/internal/app/utils"

type PriceCatcher struct {
	Date        string
	PremiseCode string
	ItemCode    string
	Price       string
}

var priceCatcherList []PriceCatcher

func init() {
	CreatePriceCatcherList(utils.ReadCsvFile("data/pricecatcher_2024-01.csv"))
}

func GetPriceCatcherList() []PriceCatcher {
	return priceCatcherList
}

func CreatePriceCatcherList(data [][]string) {
	for i, line := range data {
		if i > 0 {
			var rec PriceCatcher

			for j, field := range line {
				switch j {
				case 0:
					rec.Date = field
					break
				case 1:
					rec.PremiseCode = field
					break
				case 2:
					rec.ItemCode = field
					break
				case 3:
					rec.Price = field
					break
				default:
					break
				}
			}

			priceCatcherList = append(priceCatcherList, rec)
		}
	}
}
