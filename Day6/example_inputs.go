package main

var input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

var bigBoard = Board{
	{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
	{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", "#", ".", ".", "."}}
