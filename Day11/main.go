package main

import "fmt"

func main() {
	line := ParseInput(input)
	line.BlinkN(25)
	fmt.Printf("Num stones: %d\n", line.Len())

	sum := 0
	for stone := range line.Traverse() {
		l := []int{stone.Value}
		BlinkN(l, 50)
		sum += len(l)
		l = nil
	}
	fmt.Printf("Num stones after 75 blinks %d\n", sum)
}
