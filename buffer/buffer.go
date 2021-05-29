package buffer

import (
	"errors"
	"io"
)

const (
	defaultBufferSize = 64                 // Slice default size
	maxInt            = int(^uint(0) >> 1) // Slice max size
)

var ErrTooLarge = errors.New("buffer.Buffer: too large")

type Buffer struct {
	buf []byte
	off int
}

func NewBuffer(buf []byte) *Buffer {
	b := &Buffer{buf: buf}
	return b
}

func NewBufferString(str string) *Buffer {
	return NewBuffer([]byte(str))
}

func (b *Buffer) Len() int {
	return len(b.buf) - b.off
}

func (b *Buffer) Bytes() []byte {
	return b.buf[b.off:]
}

func (b *Buffer) Read(data []byte) (n int, err error) {
	if b.empty() {
		b.Reset()
		if len(data) == 0 {
			return 0, nil
		}
		return 0, io.EOF
	}

	n = copy(data, b.buf[b.off:])
	b.off += n
	return n, nil
}

func (b *Buffer) TryRead(data []byte, readCount func(buf []byte) int) (n int, err error) {
	if b.empty() {
		b.Reset()
		if len(data) == 0 {
			return 0, nil
		}
		return 0, io.EOF
	}

	n = readCount(b.buf[b.off:])
	copy(data, b.buf[b.off:b.off+n])
	b.off += n
	return n, nil
}

func (b *Buffer) Write(data []byte) (n int, err error) {
	m, ok := b.tryGrowByReslice(len(data))
	if !ok {
		m = b.grow(len(data))
	}

	return copy(b.buf[m:], data), nil
}

func (b *Buffer) empty() bool {
	return len(b.buf) <= b.off
}

func (b *Buffer) tryGrowByReslice(n int) (int, bool) {
	if l := len(b.buf); n <= cap(b.buf)-l {
		return l, true
	}

	return 0, false
}

func (b *Buffer) grow(n int) int {
	m := b.Len()

	if m == 0 && b.off != 0 {
		b.Reset()
	}

	if i, ok := b.tryGrowByReslice(n); ok {
		return i
	}

	if b.buf == nil && n <= defaultBufferSize {
		b.buf = make([]byte, n, defaultBufferSize)
		return 0
	}

	c := cap(b.buf)
	if n <= c/2-m {
		copy(b.buf, b.buf[b.off:])
	} else if c > maxInt-c-n {
		panic(ErrTooLarge)
	} else {
		// Not enough space anywhere, we need to allocate.
		buf := makeSlice(2*n + n)
		copy(buf, b.buf[b.off:])
		b.buf = buf
	}

	b.off = 0
	b.buf = b.buf[:m+n]
	return m
}

func (b *Buffer) Grow(n int) {
	if n < 0 {
		panic("buffer.Buffer.Grow: negative count")
	}

	m := b.grow(n)
	b.buf = b.buf[:m]
}

func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
}

// makeSlice allocates a slice of size n. If the allocation fails, it panics
// with ErrTooLarge.
func makeSlice(n int) []byte {
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()
	return make([]byte, n)
}
