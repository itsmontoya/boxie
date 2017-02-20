package boxxy

func newBlock(offset int) *block {
	return &block{offset: offset}
}

type block struct {
	buf    [32]interface{}
	tail   int // This actually could be a uint8.. hmm
	offset int
}

func (b *block) get(idx int) interface{} {
	return b.buf[idx]
}

func (b *block) append(val interface{}) (ok bool) {
	if b.tail == 32 {
		return
	}

	b.buf[b.tail] = val
	b.tail++
	return true
}

func (b *block) prepend(val interface{}) (ok bool) {
	if b.tail == 32 {
		return
	}

	b.shiftRight(0)
	b.buf[0] = val
	return true
}

func (b *block) insert(idx int, val interface{}) (overflow interface{}) {
	if b.tail == 32 {
		overflow = b.buf[31]
	}

	b.shiftRight(idx)
	b.buf[idx] = val
	return
}

func (b *block) forEach(fn func(i int, val interface{}) (end bool)) (ended bool) {
	for i := 0; i < b.tail; i++ {
		if fn(i+b.offset, b.buf[i]) {
			ended = true
			return
		}
	}

	return
}

func (b *block) shiftRight(i int) {
	var cval interface{}
	end := b.tail
	if end == 32 {
		end = 31
	} else {
		b.tail++
	}

	for ; i <= end; i++ {
		item := b.buf[i]
		b.buf[i] = cval
		cval = item
	}
}
