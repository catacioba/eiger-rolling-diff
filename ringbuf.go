package rdiff

// A RingBuffer behaves like a fixed sized Queue.
type RingBuffer interface {
	// Push adds an element to the RingBuffer. It overwrites the oldest element from the RingBuffer.
	Push(b byte)
	// Pop returns the first element in the RingBuffer. It panics if called on an empty RingBuffer.
	Pop() byte
	// Data returns the elements of the RingBuffer in their insertion order.
	Data() []byte
	Len() int
	Clear()
}

// ArrayRingBuffer is an array-backed RingBuffer implementation.
type ArrayRingBuffer struct {
	inner []byte
	start int
	end   int
	size  int
}

func NewArrayRingBuffer(size int) *ArrayRingBuffer {
	return &ArrayRingBuffer{
		inner: make([]byte, size),
		start: 0,
		end:   0,
		size:  0,
	}
}

func (a *ArrayRingBuffer) Push(b byte) {
	a.inner[a.end] = b
	a.end++
	a.end %= cap(a.inner)

	if a.size == cap(a.inner) {
		a.start++
	} else {
		a.size++
	}
}

func (a *ArrayRingBuffer) Pop() byte {
	if a.size == 0 {
		panic("Pop() on empty ring buffer")
	}
	first := a.inner[a.start]
	a.start++
	a.start %= cap(a.inner)
	a.size--
	return first
}

func (a *ArrayRingBuffer) Data() []byte {
	idx := 0
	data := make([]byte, a.size)

	if a.size == 0 {
		return data
	}
	if a.start < a.end {
		for i := a.start; i < a.end; i++ {
			data[idx] = a.inner[i]
			idx++
		}
	} else {
		for i := a.start; i < cap(a.inner); i++ {
			data[idx] = a.inner[i]
			idx++
		}
		for i := 0; i < a.end; i++ {
			data[idx] = a.inner[i]
			idx++
		}
	}

	return data
}

func (a *ArrayRingBuffer) Len() int {
	return a.size
}

func (a *ArrayRingBuffer) Clear() {
	a.start = 0
	a.end = 0
	a.size = 0
}
