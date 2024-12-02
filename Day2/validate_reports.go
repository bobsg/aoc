package main

const directionAsc = 0
const directionDesc = 1

func ValidateReports(input [][]int) int {
	var numSafe int
	for _, row := range input {
		currentDirection := -1
		lastNum := row[0]
		safe := true
		for i := 1; i < len(row); i++ {
			var direction int
			if i == 1 {
				if row[0] > row[i] {
					currentDirection = directionDesc
				} else {

				}
			}
		}
	}
}
