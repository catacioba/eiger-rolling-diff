package rdiff

import "crypto/md5"

// MOD is the largest prime number that fits in an uint16.
const MOD = 65521

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
	r.bn *= 256
	r.bn %= MOD
	r.sum *= 256
	r.sum %= MOD
	r.sum += r.bn * uint32(b)
	r.sum %= MOD
}

func (r *PolynomialRollingHash) RotatePush(in byte, out byte) {
	r.sum -= r.bn * uint32(out)
	r.sum %= MOD
	r.sum *= 256
	r.sum %= MOD
	r.sum += uint32(in)
	r.sum %= MOD
}

func (r *PolynomialRollingHash) ChuckSum() uint16 {
	return uint16(r.sum & 0xFFFF)
}

func StrongCheckSum(data []byte) [16]byte {
	return md5.Sum(data)
}
