package main

import (
	"log"
	"strconv"
	"strings"
)

func ParseInput(input string) *Line {
	parts := strings.Split(input, " ")
	l := &Line{}
	for _, v := range parts {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Panicf("could not parse %s", v)
		}
		l.Add(val)
	}
	return l
}
