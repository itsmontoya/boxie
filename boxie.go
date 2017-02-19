package boxie

func New() *Boxie {
	var b Boxie
	b.bs = append(b.bs, newBlock())
	return &b
}

type Boxie struct {
	bs   []*block
	tail int
}

func (b *Boxie) getContainingBucketIdx(idx int) (i, rem int) {
	var ci int
	for bi, b := range b.bs {
		if idx >= ci && idx < ci+b.tail {
			i = bi
			rem = idx - ci
			return
		}

		ci += b.tail
	}

	return -1, 0
}

func (b *Boxie) shiftRight(idx int) {
	var cb *block
	for i, item := range b.bs[idx:] {
		b.bs[i] = cb
		cb = item
	}

	b.bs = append(b.bs, cb)
	b.tail++
}

// Get will get an item at a provided index
func (b *Boxie) Get(idx int) interface{} {
	bi, rem := b.getContainingBucketIdx(idx)

	if bi == -1 {
		return nil
	}

	blk := b.bs[bi]
	return blk.get(rem)
}

// Append will append an item to the list
func (b *Boxie) Append(val interface{}) {
	lb := b.bs[b.tail]
	if lb == nil || !lb.append(val) {
		lb = newBlock()
		lb.append(val)
		b.bs = append(b.bs, lb)
		b.tail++
	}
}

// Prepend will insert an item to the beginning of the list
func (b *Boxie) Prepend(val interface{}) {
	fb := b.bs[0]
	if fb == nil || !fb.prepend(val) {
		fb = newBlock()
		// Since this is our first item, we know we can append (it's faster)
		fb.append(val)
		b.shiftRight(0)
		b.bs[0] = fb
	}
}

// Insert will place an item at the requested index within a list
func (b *Boxie) Insert(idx int, val interface{}) {
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
		blk = newBlock()
		blk.append(of)
		b.bs = append(b.bs, blk)
		b.tail++
		return
	}

	blk = b.bs[bi]
	if blk.prepend(of) {
		return
	}

	blk = newBlock()
	blk.append(of)
	b.shiftRight(bi)
	b.bs[bi] = blk
}

// ForEach will iterate through each item within Boxie
func (b *Boxie) ForEach(fn func(i int, val interface{}) (end bool)) (ended bool) {
	var offset int
	for _, b := range b.bs {
		if b.forEach(offset, fn) {
			ended = true
			return
		}

		offset += b.tail
	}

	return
}

func newBlock() *block {
	return &block{}
}

type block struct {
	buf  [32]interface{}
	tail int
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

func (b *block) forEach(offset int, fn func(i int, val interface{}) (end bool)) (ended bool) {
	for i := 0; i < b.tail; i++ {
		if fn(i+offset, b.buf[i]) {
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
