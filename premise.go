package main

type Premise struct {
	PremiseCode string
	Premise     string
	Address     string
	PremiseType string
	State       string
	District    string
}

func createPremiseList(data [][]string) []Premise {
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

