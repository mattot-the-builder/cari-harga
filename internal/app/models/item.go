package models

type Item struct {
	ItemCode          string
	ItemName          string
	UnitOfMeasurement string
	ItemGroup         string
	ItemCategory      string
}

func CreateItemList(data [][]string) []Item {
	var itemList []Item
	for i, line := range data {
		if i > 0 {
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
	return itemList
}
