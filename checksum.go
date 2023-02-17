package rdiff

import "crypto/md5"

// MOD is the largest prime number that fits in an uint16.
const MOD = 65521

const BASE = 256

// PolynomialRollingHash computes a rolling hash using a simple polynomial algorithm.
//
// The calling code is responsible for correctly rotating the values (i.e. it should keep track of
// the correct order of values).
type PolynomialRollingHash struct {
	sum uint32
	bn  uint32
}

func NewPolynomialRollingHash() PolynomialRollingHash {
	return PolynomialRollingHash{
		sum: uint32(0),
		bn:  uint32(1),
	}
}

func (r *PolynomialRollingHash) Push(b byte) {
	r.bn *= BASE
	r.bn %= MOD
	r.sum *= BASE
	r.sum %= MOD
	r.sum += uint32(b)
	r.sum %= MOD
}

func (r *PolynomialRollingHash) RotatePush(in byte, out byte) {
	r.sum *= BASE
	r.sum %= MOD
	r.sum -= (r.bn * uint32(out)) % MOD
	r.sum += MOD // to make sure we are working with positive remainders.
	r.sum %= MOD
	r.sum += uint32(in)
	r.sum %= MOD
}

func (r *PolynomialRollingHash) CheckSum() uint16 {
	return uint16(r.sum & 0xFFFF)
}

func (r *PolynomialRollingHash) Reset() {
	r.sum = uint32(0)
	r.bn = uint32(1)
}

func WeakCheckSum(data []byte) uint16 {
	h := NewPolynomialRollingHash()
	for _, d := range data {
		h.Push(d)
	}
	return h.CheckSum()
}

func StrongCheckSum(data []byte) [16]byte {
	return md5.Sum(data)
}
