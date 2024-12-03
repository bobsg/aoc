package main

func RunMemory(memory [][]int) int {
	var sum int
	for _, p := range memory {
		sum += p[0] * p[1]
	}
	return sum
}

func RunMemoryWithConditionals(memory []Command) int {
	var sum int
	var active = true
	for _, cmd := range memory {
		switch cmd.Type {
		case MUL:
			if active {
				sum += cmd.X * cmd.Y
			}
		case DO:
			active = true
		case DONT:
			active = false
		}

	}

	return sum
}
