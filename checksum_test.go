package rdiff

import (
	"testing"
)

func TestFullyRotatedSum(t *testing.T) {
	h := NewPolynomialRollingHash()

	h.Push(10)
	h.Push(20)
	h.Push(30)
	h.Push(40)
	sum1 := h.CheckSum()

	h.RotatePush(10, 10)
	h.RotatePush(20, 20)
	h.RotatePush(30, 30)
	h.RotatePush(40, 40)
	sum2 := h.CheckSum()

	if sum1 != sum2 {
		t.Errorf("Fully rotated sums should be equal")
	}
}

func TestRotatePush(t *testing.T) {
	h := NewPolynomialRollingHash()

	h.Push(10)
	h.Push(20)
	h.Push(30)

	h.RotatePush(40, 10)

	if h.CheckSum() != 8020 {
		t.Errorf("Incorrect checksum")
	}

	h.RotatePush(50, 20)

	if h.CheckSum() != 10740 {
		t.Errorf("Incorrect checksum")
	}
}

func TestWeakCheckSum(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want uint16
	}{
		{
			name: "one element",
			data: []byte{10},
			want: 10,
		},
		{
			name: "two elements",
			data: []byte{10, 20},
			want: 2580,
		},
		{
			name: "three elements",
			data: []byte{10, 20, 30},
			want: 5300,
		},
		{
			name: "four elements",
			data: []byte{10, 20, 30, 40},
			want: 46420,
		},
		{
			name: "five elements",
			data: []byte{10, 20, 30, 40, 50},
			want: 24269,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeakCheckSum(tt.data); got != tt.want {
				t.Errorf("WeakCheckSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
