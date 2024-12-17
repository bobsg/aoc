package main

import (
	"fmt"
	"iter"
	"log"
	"strconv"
	"strings"
)

type Line struct {
	length int
	First  *Stone
	Last   *Stone
}

type Line2 []int

type Stone struct {
	Value int
	Next  *Stone
	Prev  *Stone
}

func (l *Line) Len() int {
	return l.length
}

func (l *Line) Add(val int) {
	s := &Stone{Value: val}
	if l.Last == nil {
		l.First = s
	} else {
		s.Prev = l.Last
		l.Last.Next = s
	}
	l.Last = s
	l.length++
}

func (l *Line) Split(s *Stone) {
	curr := s
	if curr == nil {
		return
	}

	value := strconv.Itoa(curr.Value)
	i := len(value) / 2
	part1, err := strconv.Atoi(value[:i])
	if err != nil {
		log.Panicf("could not convert part 1 of %s to int", value)
	}
	part2, err := strconv.Atoi(value[i:])
	if err != nil {
		log.Panicf("could not convert part 2 of %s to int", value)
	}

	s.Value = part1
	s2 := &Stone{Value: part2, Prev: s, Next: s.Next}
	if l.Last == s {
		l.Last = s2
	} else {
		s.Next.Prev = s2
	}
	s2.Prev = s
	s.Next = s2
	l.length++

}

func (l *Line) Delete(s *Stone) {
	if s.Next != nil {
		s.Next.Prev = s.Prev
	}
	if s.Prev != nil {
		s.Prev.Next = s.Next
	}
	if l.First == s {
		l.First = s.Next
	}
	if l.Last == s {
		l.Last = s.Prev
	}
	l.length--
}

func (l *Line) FindStone(s *Stone) *Stone {
	curr := l.First

	for curr != nil && curr != s {
		curr = curr.Next
	}

	return curr
}

func (l *Line) String() string {
	var b strings.Builder
	curr := l.First
	for curr != nil {
		_, _ = fmt.Fprintf(&b, "%d", curr.Value)
		if curr != l.Last {
			_, _ = fmt.Fprint(&b, " ")
		}
		curr = curr.Next
	}

	return b.String()
}

func String(l []int) string {
	var b strings.Builder
	length := len(l) - 1
	for i, n := range l {
		_, _ = fmt.Fprintf(&b, "%d", n)
		if i != length {
			_, _ = fmt.Fprint(&b, " ")
		}
	}

	return b.String()
}

func (l *Line) StringRev() string {
	var b strings.Builder
	curr := l.Last
	for curr != nil {
		_, _ = fmt.Fprintf(&b, "%d", curr.Value)
		if curr != l.First {
			_, _ = fmt.Fprint(&b, " ")
		}
		curr = curr.Prev
	}

	return b.String()
}
func (l *Line) BlinkNPreCalc(n int) int {
	for i := 0; i < min(n, 30); i++ {
		l.Blink()
	}
	if min(n, 30) == n {
		return l.Len()
	}
	fmt.Println("Starting precalc")
	preCalc := Fill()
	fmt.Println("Precalc done")
	sum := 0
	for i := 30; i < n; i++ {
		curr := l.First
		for curr != nil {
			if curr.Value < 10 {
				sum += preCalc[curr.Value][n-i]
				l.Delete(curr)
			} else if len(strconv.Itoa(curr.Value))%2 == 0 {
				l.Split(curr)
				curr = curr.Next // jump over newly created stone
			} else {
				curr.Value *= 2024
			}
			curr = curr.Next
		}
	}
	sum += l.Len()
	return sum
}

func Fill() []map[int]int {
	result := make([]map[int]int, 10)

	for i := 0; i <= 9; i++ {
		result[i] = map[int]int{}
		l := Line{}
		l.Add(i)
		for x := 1; x <= 45; x++ {
			l.Blink()
			result[i][x] = l.Len()
		}
	}
	return result
}

func (l *Line) Blink() {
	curr := l.First
	for curr != nil {
		if curr.Value == 0 {
			curr.Value = 1
		} else if len(strconv.Itoa(curr.Value))%2 == 0 {
			l.Split(curr)
			curr = curr.Next // jump over newly created stone
		} else {
			curr.Value *= 2024
		}
		curr = curr.Next
	}
}

func Blink(l []int) []int {
	length := len(l)
	for i := 0; i < length; i++ {
		if l[i] == 0 {
			l[i] = 1
		} else if len(strconv.Itoa(l[i]))%2 == 0 {
			a, b := Split(l[i])
			l[i] = a
			l = append(l, b)
		} else {
			l[i] *= 2024
		}
	}
	return l
}

func Split(n int) (int, int) {
	value := strconv.Itoa(n)
	i := len(value) / 2
	part1, err := strconv.Atoi(value[:i])
	if err != nil {
		log.Panicf("could not convert part 1 of %s to int", value)
	}
	part2, err := strconv.Atoi(value[i:])
	if err != nil {
		log.Panicf("could not convert part 2 of %s to int", value)
	}

	return part1, part2

}

func BlinkN(l []int, n int) []int {
	for i := 0; i < n; i++ {
		l = Blink(l)
	}
	return l
}

func (l *Line) BlinkN(n int) {
	for i := 0; i < n; i++ {
		l.Blink()
	}
}

func (l *Line) Traverse() iter.Seq[*Stone] {
	return func(yield func(*Stone) bool) {
		curr := l.First
		for curr != nil {
			if !yield(curr) {
				return
			}
			curr = curr.Next
		}
	}

}
