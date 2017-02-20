package boxxy

import (
//	"fmt"
)

// New will return a new instance of Boxxie
func New() *Boxxy {
	var b Boxxy
	b.bs = append(b.bs, newBlock(0))
	return &b
}

// Boxxy is a sharded slice
type Boxxy struct {
	bs   []*block
	tail int
}

func (b *Boxxy) getContainingBucketIdx(idx int) (i, rem int) {
	guess := idx / 32
	if guess > b.tail {
		guess = b.tail
	}

	for {
		gos := b.bs[guess].offset
		gt := b.bs[guess].tail
		if gos > idx {
			guess--
			continue
		}

		if gos+gt < idx {
			guess++
			continue
		}

		i = guess
		rem = idx - gos
		return
	}

	return -1, 0
}

func (b *Boxxy) shiftRight(idx int) {
	var cb *block
	for i, item := range b.bs[idx:] {
		b.bs[i] = cb
		cb = item
	}

	b.bs = append(b.bs, cb)
	b.tail++
}

// Get will get an item at a provided index
func (b *Boxxy) Get(idx int) interface{} {
	bi, rem := b.getContainingBucketIdx(idx)

	if bi == -1 {
		return nil
	}

	blk := b.bs[bi]
	return blk.get(rem)
}

// Append will append an item to the list
func (b *Boxxy) Append(val interface{}) {
	lb := b.bs[b.tail]
	if lb == nil || !lb.append(val) {
		// Last block's offset
		po := b.bs[b.tail].offset
		// Last block's tail
		pt := b.bs[b.tail].tail

		lb = newBlock(po + pt)
		lb.append(val)

		b.bs = append(b.bs, lb)
		b.tail++
	}
}

func (b *Boxxy) incrementOffset(start int) {
	for i := start; i < len(b.bs); i++ {
		b.bs[i].offset++
	}
}

// Prepend will insert an item to the beginning of the list
func (b *Boxxy) Prepend(val interface{}) {
	fb := b.bs[0]
	if fb == nil || !fb.prepend(val) {
		fb = newBlock(0)
		// Since this is our first item, we know we can append (it's faster)
		fb.append(val)
		b.shiftRight(0)
		b.bs[0] = fb
	}

	b.incrementOffset(1)
}

// Insert will place an item at the requested index within a list
func (b *Boxxy) Insert(idx int, val interface{}) {
	bi, rem := b.getContainingBucketIdx(idx)
	if bi == -1 {
		// ERROR OUT OF RANGE
		return
	}

	blk := b.bs[bi]
	of := blk.insert(rem, val)
	if of == nil {
		return
	}

	bi++
	if bi > b.tail {
		blk = newBlock(idx)
		blk.append(of)
		b.bs = append(b.bs, blk)
		b.tail++
		goto END
	}

	blk = b.bs[bi]
	if blk.prepend(of) {
		goto END
	}

	blk = newBlock(idx)
	blk.append(of)
	b.shiftRight(bi)
	b.bs[bi] = blk

END:
	b.incrementOffset(idx + 1)
}

// ForEach will iterate through each item within Boxxy
func (b *Boxxy) ForEach(fn func(i int, val interface{}) (end bool)) (ended bool) {
	for _, b := range b.bs {
		if b.forEach(fn) {
			ended = true
			return
		}
	}

	return
}
