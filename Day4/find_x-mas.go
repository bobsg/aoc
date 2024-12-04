package main

func FindA(input [][]string) []Pos {
	var ret []Pos
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "A" {
				p := Pos{X: j, Y: i}
				ret = append(ret, p)
			}
		}
	}

	return ret
}

func isInBounds2(input [][]string, p Pos) bool {
	if p.X == 0 || p.X == len(input[p.Y])-1 {
		return false
	}

	if p.Y == 0 || p.Y == len(input)-1 {
		return false
	}

	return true
}

func CountCrossXmases(input [][]string) int {
	var sum int
	as := FindA(input)
	for _, a := range as {
		sum += testPosition(input, a)
	}

	return sum
}

func testPosition(input [][]string, p Pos) int {
	if !isInBounds2(input, p) {
		return 0
	}

	if checkFrontslash(input, p) && checkBackslash(input, p) {
		return 1
	}

	return 0
}

func checkBackslash(input [][]string, p Pos) bool {
	ul := Pos{p.X - 1, p.Y - 1}
	lr := Pos{p.X + 1, p.Y + 1}

	if input[ul.Y][ul.X] == "M" {
		return input[lr.Y][lr.X] == "S"
	} else if input[ul.Y][ul.X] == "S" {
		return input[lr.Y][lr.X] == "M"
	}

	return false
}

func checkFrontslash(input [][]string, p Pos) bool {
	ur := Pos{p.X + 1, p.Y - 1}
	ll := Pos{p.X - 1, p.Y + 1}

	if input[ur.Y][ur.X] == "M" {
		return input[ll.Y][ll.X] == "S"
	} else if input[ur.Y][ur.X] == "S" {
		return input[ll.Y][ll.X] == "M"
	}

	return false
}
