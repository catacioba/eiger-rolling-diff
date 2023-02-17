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

	expectedData := []byte{1}
	if !bytes.Equal(ringBuf.Data(), expectedData) {
		t.Errorf("Data() should be %v", expectedData)
	}

	ringBuf.Push(2)
	ringBuf.Push(3)
	v := ringBuf.Pop()

	expectedPop := byte(1)
	if v != expectedPop {
		t.Errorf("Pop() should return %d", expectedPop)
	}
	expectedData = []byte{2, 3}
	if !bytes.Equal(ringBuf.Data(), expectedData) {
		t.Errorf("Data() should be %v", expectedData)
	}

	ringBuf.Push(4)

	expectedData = []byte{2, 3, 4}
	if !bytes.Equal(ringBuf.Data(), expectedData) {
		t.Errorf("Data() should be %v", expectedData)
	}

	ringBuf.Push(5)

	expectedData = []byte{3, 4, 5}
	if !bytes.Equal(ringBuf.Data(), expectedData) {
		t.Errorf("Data() should be %v", expectedData)
	}

	expectedPop = 3
	if ringBuf.Pop() != expectedPop {
		t.Errorf("Pop() should return %d", expectedPop)
	}
	expectedData = []byte{4, 5}
	if !bytes.Equal(ringBuf.Data(), expectedData) {
		t.Errorf("Data() should be %v", expectedData)
	}

	expectedPop = 4
	if ringBuf.Pop() != expectedPop {
		t.Errorf("Pop() should return %d", expectedPop)
	}
	expectedData = []byte{5}
	if !bytes.Equal(ringBuf.Data(), expectedData) {
		t.Errorf("Data() should be %v", expectedData)
	}

	expectedPop = 5
	if ringBuf.Pop() != expectedPop {
		t.Errorf("Pop() should return %d", expectedPop)
	}
	expectedData = []byte{}
	if !bytes.Equal(ringBuf.Data(), expectedData) {
		t.Errorf("Data() should be %v", expectedData)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pop() on empty ring buffer should panic")
		}
	}()
	ringBuf.Pop()
}
