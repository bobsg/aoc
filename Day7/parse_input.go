package main

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

func ParseInput(r io.Reader) []Equation {
	result := []Equation{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == " " {
			continue
		}
		parts := strings.Split(line, ":")
		valueText := strings.Trim(parts[0], " ")
		value, err := strconv.Atoi(valueText)
		if err != nil {
			log.Fatalf("could not parse %s\n", valueText)
		}
		numbersText := strings.Split(strings.Trim(parts[1], " "), " ")
		numbers := []int{}
		for i := 0; i < len(numbersText); i++ {
			n, err := strconv.Atoi(numbersText[i])
			if err != nil {
				log.Fatalf("could not parse %s\n", numbersText[i])
			}
			numbers = append(numbers, n)
		}
		e := Equation{
			TestValue: value,
			Numbers:   numbers,
		}
		result = append(result, e)
	}

	return result
}
