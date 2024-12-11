package main

import (
	"bufio"
	"io"
)

func ParseInput(reader io.Reader) Antennas {
	result := map[frequency][]Antenna{}

	scanner := bufio.NewScanner(reader)
	height := 0
	width := 0
	for scanner.Scan() {
		line := scanner.Text()
		if width == 0 {
			width = len(line)
		}
		for x, r := range []rune(line) {
			freq := frequency(r)
			if freq != '.' {
				result[freq] = append(result[freq], Antenna{freq, Position{X: x, Y: height}})
			}
		}
		height++
	}
	return Antennas{
		width:    width,
		height:   height,
		antennas: result,
	}
}
