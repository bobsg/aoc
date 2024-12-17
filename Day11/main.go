package main

import "fmt"

func main() {

	line := ParseInput(input)
	line.BlinkN(25)
	fmt.Printf("Num stones: %d\n", line.Len())

	l2 := ParseInput(input)
	sum := l2.BlinkNPreCalc(75)
	fmt.Printf("Num stones2: %d\n", sum)

}
