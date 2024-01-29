package models

import (
	"strings"

	"github.com/mattot-the-builder/go-csv/internal/app/utils"
)

type Item struct {
	ItemCode          string
	ItemName          string
	UnitOfMeasurement string
	ItemGroup         string
	ItemCategory      string
}

var itemList []Item

func init() {
	CreateItemList(utils.ReadCsvFile("data/lookup_item.csv"))
}

func GetItemList() []Item {
	return itemList
}

func SearchItem(keyword string) []Item {
	var searchResult []Item
	for _, item := range itemList {
		if strings.Contains(item.ItemName, strings.ToUpper(keyword)) {
			searchResult = append(searchResult, item)
		}
	}
	return searchResult
}

func CreateItemList(data [][]string) {
	for i, line := range data {
		if i > 1 {
			var rec Item

			for j, field := range line {
				switch j {
				case 0:
					rec.ItemCode = field
					break
				case 1:
					rec.ItemName = field
					break
				case 2:
					rec.UnitOfMeasurement = field
					break
				case 3:
					rec.ItemGroup = field
					break
				case 4:
					rec.ItemCategory = field
					break
				default:
					break
				}
			}

			itemList = append(itemList, rec)
		}

	}
}
