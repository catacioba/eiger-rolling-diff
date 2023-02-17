package rdiff

import (
	"bytes"
	"testing"
)

func TestArrayRingBuffer(t *testing.T) {
	ringBuf := NewArrayRingBuffer(3)

	if !bytes.Equal(ringBuf.Data(), []byte{}) {
		t.Errorf("Data() should be empty")
	}

	ringBuf.Push(1)

	if !bytes.Equal(ringBuf.Data(), []byte{1}) {
		t.Errorf("Data() should have only 1 element")
	}

	ringBuf.Push(2)
	ringBuf.Push(3)
	v := ringBuf.Pop()

	if v != 1 {
		t.Errorf("Pop() should return 1")
	}
	if !bytes.Equal(ringBuf.Data(), []byte{2, 3}) {
		t.Errorf("Data() should be [2, 3]")
	}

	ringBuf.Push(4)

	if !bytes.Equal(ringBuf.Data(), []byte{2, 3, 4}) {
		t.Errorf("Data() should be [2, 3, 4]")
	}

	ringBuf.Push(5)

	if !bytes.Equal(ringBuf.Data(), []byte{3, 4, 5}) {
		t.Errorf("Data() should be [3, 4, 5]")
	}

	if ringBuf.Pop() != 3 {
		t.Errorf("Pop() should return 3")
	}
	if !bytes.Equal(ringBuf.Data(), []byte{4, 5}) {
		t.Errorf("Data() should be [4, 5]")
	}

	if ringBuf.Pop() != 4 {
		t.Errorf("Pop() should return 4")
	}
	if !bytes.Equal(ringBuf.Data(), []byte{5}) {
		t.Errorf("Data() should be [5]")
	}

	if ringBuf.Pop() != 5 {
		t.Errorf("Pop() should return 5")
	}
	if !bytes.Equal(ringBuf.Data(), []byte{}) {
		t.Errorf("Data() should be []")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pop() on empty ring buffer should panic")
		}
	}()
	ringBuf.Pop()
}
