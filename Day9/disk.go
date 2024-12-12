package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Disk struct {
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

func NewDisk() *Disk {
	return &Disk{}
}

func (d *Disk) AddBlock(id, size int) {
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

func (d *Disk) AddFreeBlock(size int) {
	d.AddBlock(-1, size)
}

func (d *Disk) Len() int {
	return d.Length
}

func (d *Disk) String() string {
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

func (d *Disk) FirstFree() *Block {
	curr := d.First
	for curr != nil && curr.Id != -1 {
		curr = curr.Next
	}

	return curr
}

func (d *Disk) FindBlockById(id int) *Block {
	curr := d.First
	for curr != nil {
		curr = curr.Next
		if curr.Id == id {
			break
		}
	}
	return curr
}

func (d *Disk) InsertFreeAfter(size int, b *Block) {
	d.InsertAfter(-1, size, b)
}

func (d *Disk) InsertFreeBefore(size int, b *Block) {
	d.InsertBefore(-1, size, b)
}

func (d *Disk) InsertAfter(id, size int, b *Block) {
	block := &Block{Id: id, Size: size}
	if d.Last == b {
		d.Last = block
	}
	block.Next = b.Next
	block.Previous = b
	b.Next = block
	d.Length++
}

func (d *Disk) InsertBefore(id, size int, b *Block) {
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

func (d *Disk) Delete(b *Block) {
	if b.Previous != nil {
		b.Previous.Next = b.Next
	}
	if b.Next != nil {
		b.Next.Previous = b.Previous
	}

	if d.First == b {
		d.First = b.Next
	}
	if d.Last == b {
		d.Last = b.Previous
	}
	d.Length--
	//d.TryMerge(b.Previous, b.Next)
}

func (d *Disk) TryMerge(a, b *Block) {
	if a == nil || b == nil {
		return
	}
	if a.Id != b.Id {
		return
	}

	a.Size += b.Size
	b.Size = 0
	d.Delete(b)
}

func (d *Disk) Shrink(size int, b *Block) {
	b.Size--
	if b.Size == 0 {
		d.Delete(b)
	}
}

func (d *Disk) Defrag() {
	// The first block is never a free block
	// Loop until first free block is the same as the disks last block
	firstFree := d.FirstFree()
	for firstFree != d.Last {
		blockToDefrag := getLastNonFree(d.Last)
		id := blockToDefrag.Id

		// Remove at end and increase free block after or add new one
		if blockToDefrag.Next == nil {
			d.InsertFreeAfter(1, blockToDefrag)
		} else {
			blockToDefrag.Next.Size++
		}
		d.Shrink(1, blockToDefrag)

		// Add before first free block
		if firstFree.Previous != nil && firstFree.Previous.Id == id {
			firstFree.Previous.Size++
		} else {
			d.InsertBefore(id, 1, firstFree)
		}

		d.Shrink(1, firstFree)

		firstFree = d.FirstFree()
	}
}

func getLastNonFree(b *Block) *Block {
	curr := b
	for curr != nil && curr.Id == -1 {
		curr = curr.Previous
	}
	return curr
}
