package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DiskP2 struct {
	Length int
	First  *Block
	Last   *Block
}

type Block struct {
	Id       int
	Size     int
	Previous *Block
	Next     *Block
}

func NewDiskP2() *DiskP2 {
	return &DiskP2{}
}

func (d *DiskP2) AddBlock(id, size int) {
	if size == 0 {
		return
	}

	b := &Block{Id: id, Size: size}
	if d.Last == nil {
		d.First = b
	} else {
		b.Previous = d.Last
		d.Last.Next = b
	}
	d.Last = b
	d.Length++
}

func (d *DiskP2) AddFreeBlock(size int) {
	d.AddBlock(-1, size)
}

func (d *DiskP2) Len() int {
	return d.Length
}

func (d *DiskP2) String() string {
	var b strings.Builder
	curr := d.First
	for curr != nil {
		s := "."
		if curr.Id != -1 {
			s = strconv.Itoa(curr.Id)
		}
		for i := 0; i < curr.Size; i++ {
			_, _ = fmt.Fprintf(&b, "%s", s)
		}
		curr = curr.Next
	}

	return b.String()
}

func (d *DiskP2) FirstFree(size int) *Block {
	curr := d.First
	for curr != nil {
		if curr.Id == -1 && curr.Size >= size {
			return curr
		}
		curr = curr.Next
	}

	return nil
}

func (d *DiskP2) FindBlockById(id int) *Block {
	curr := d.First
	for curr != nil {
		curr = curr.Next
		if curr.Id == id {
			break
		}
	}
	return curr
}

func (d *DiskP2) InsertFreeAfter(size int, b *Block) {
	d.InsertAfter(-1, size, b)
}

func (d *DiskP2) InsertFreeBefore(size int, b *Block) {
	d.InsertBefore(-1, size, b)
}

func (d *DiskP2) InsertAfter(id, size int, b *Block) {
	block := &Block{Id: id, Size: size}
	if d.Last == b {
		d.Last = block
	}
	block.Next = b.Next
	block.Previous = b
	b.Next = block
	d.Length++
}

func (d *DiskP2) InsertBefore(id, size int, b *Block) {
	block := &Block{Id: id, Size: size}
	block.Next = b
	block.Previous = b.Previous
	b.Previous = block
	if block.Previous != nil {
		block.Previous.Next = block
	}
	if b == d.First {
		d.First = block
	}
	d.Length++
}

func (d *DiskP2) Defrag() {
	curr := d.Last
	for curr != nil {
		nextInIteration := curr.Previous
		if curr.Id != -1 {
			b := d.FirstFreeBeforeCurrent(curr)
			if b != nil {
				d.MoveToFree(curr, b)
			}
		}
		curr = nextInIteration
	}
}

func (d *DiskP2) FirstFreeBeforeCurrent(block *Block) *Block {
	curr := d.First
	for curr != nil && curr != block {
		if curr.Id == -1 && curr.Size >= block.Size {
			return curr
		}
		curr = curr.Next
	}

	return nil
}

func (d *DiskP2) MoveToFree(block, freeBlock *Block) {
	if freeBlock.Size < block.Size {
		log.Panicf("block with id %d and size %d cannot fit in free block with size %d", block.Id, block.Size, freeBlock.Size)
	}

	if freeBlock.Size == block.Size {
		freeBlock.Id = block.Id
		block.Id = -1
		return
	}

	freeBlock.Size -= block.Size
	d.InsertBefore(block.Id, block.Size, freeBlock)
	block.Id = -1
}

func getLastNonFree(b *Block) *Block {
	curr := b
	for curr != nil && curr.Id == -1 {
		curr = curr.Previous
	}
	return curr
}

func (d *DiskP2) Checksum() int {
	blocks := d.ToSlice()
	sum := 0
	for i, v := range blocks {
		if v != -1 {
			sum += i * v
		}
	}

	return sum
}

func (d *DiskP2) ToSlice() []int {
	s := []int{}
	curr := d.First
	for curr != nil {
		for i := 0; i < curr.Size; i++ {
			s = append(s, curr.Id)
		}
		curr = curr.Next
	}
	return s
}
