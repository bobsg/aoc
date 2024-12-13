package main

import (
	"fmt"
	"strings"
)

type Disk struct {
	Blocks []int
}

func (d *Disk) AddBlock(id int) {
	d.Blocks = append(d.Blocks, id)
}

func (d *Disk) AddBlocks(id, n int) {
	for i := 0; i < n; i++ {
		d.AddBlock(id)
	}
}

func NewDisk() *Disk {
	return &Disk{
		Blocks: []int{},
	}
}

func (d *Disk) String() string {
	var b strings.Builder
	for i := 0; i < len(d.Blocks); i++ {
		if d.Blocks[i] == -1 {
			_, _ = fmt.Fprintf(&b, "%s", ".")
		} else {
			_, _ = fmt.Fprintf(&b, "%d", d.Blocks[i])
		}
	}
	return b.String()
}

func (d *Disk) Defrag() {
	// Find first free and last non free and swap them
	ff := d.firstFree(0)
	last := d.lastNonFree(len(d.Blocks) - 1)

	for ff < last {
		d.Blocks[ff], d.Blocks[last] = d.Blocks[last], d.Blocks[ff]
		ff = d.firstFree(ff)
		last = d.lastNonFree(last)
	}

}

func (d *Disk) firstFree(start int) int {
	for i := start; i < len(d.Blocks); i++ {
		if d.Blocks[i] == -1 {
			return i
		}
	}

	return -1
}

func (d *Disk) lastNonFree(start int) int {
	for i := start; i >= 0; i-- {
		if d.Blocks[i] != -1 {
			return i
		}
	}

	return -1
}

func (d *Disk) Checksum() int {
	sum := 0
	for i, v := range d.Blocks {
		if v != -1 {
			sum += i * v
		}
	}

	return sum
}
