package main

var tinyExample = `AA
BB`

var shortExample = `AAAA
BBCD
BBCC
EEEC`

var regionInRegion = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

var largerExample = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

var shortExampleParsed = [][]string{{"A", "A", "A", "A"},
	{"B", "B", "C", "D"},
	{"B", "B", "C", "C"},
	{"E", "E", "E", "C"}}
