// Package bitset implements getting and setting of true/false values within bytes.
package bitset

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// According to the Go spec, byte = uint8, so we set our byte width to 8 here
const (
	WIDTH = 8
)

// BitSet is simply an array of bytes large enough to contain the number of bits the
// user requests. Values are encoded in a little endian fashion: in a single bytes, bits
// would be indexes as 7 | 6 | 5 | 4 | 3 | 2 | 1 | 0. Callers need not concern themselves
// with this implementation detail: getting, setting and unsetting bits is done through
// functions.
type BitSet struct {
	bytesArr []byte
	size     int
}

// NewBitSet will create a new BitSet that can accomodate n bits.
func NewBitSet(n uint32) *BitSet {
	bs := new(BitSet)
	if n == 0 {
		bs.bytesArr = make([]byte, 0)
	} else {
		// round up so that the number of bytes can accomodate n bits
		bs.bytesArr = make([]byte, ((n-1)/WIDTH)+1)
	}
	bs.size = len(bs.bytesArr) * WIDTH
	return bs
}

// New is an alias for NewBitSet.
func New(n uint32) *BitSet {
	return NewBitSet(n)
}

// Size returns the size of the BitSet.
func (bs *BitSet) Size() int {
	return bs.size
}

// DumpBitSet will return a string form of the BitSet.
func (bs *BitSet) DumpBitSet() string {
	pss := ""
	for _, v := range bs.bytesArr {
		s := strconv.FormatInt(int64(v), 2)
		d := WIDTH - len(s)
		ps := ""
		for i := 0; i < d; i++ {
			ps += "0"
		}
		pss += " " + ps + s
	}
	return strings.TrimSpace(pss)
}

// String satisifes the stringer interface.
func (bs *BitSet) String() string {
	return bs.DumpBitSet()
}

// Calculate byte and bit offsets for operations.
// Error if n is not indexable in this BitSet.
func (bs *BitSet) offsets(n int) (int, byte, error) {
	if n < 0 || n > (len(bs.bytesArr)*WIDTH) {
		e := fmt.Sprintf("index %d under or overflows bitset", n)
		return 0, 0, errors.New(e)
	}
	byteOffset := len(bs.bytesArr) - (n / WIDTH) - 1
	bitOffset := byte(n % WIDTH)
	return byteOffset, bitOffset, nil
}

// SetBitN will set bit n to be 1.
// Error if n is not indexable in this BitSet.
func (bs *BitSet) SetBitN(n int) error {
	byteOffset, bitOffset, offsetErr := bs.offsets(n)
	if offsetErr != nil {
		return offsetErr
	}
	bs.bytesArr[byteOffset] |= (1 << bitOffset)
	return nil
}

// UnsetBitN will set bit n to be 0.
// Error if n is not indexable in this BitSet.
func (bs *BitSet) UnsetBitN(n int) error {
	byteOffset, bitOffset, offsetErr := bs.offsets(n)
	if offsetErr != nil {
		return offsetErr
	}
	bs.bytesArr[byteOffset] ^= (1 << bitOffset)
	return nil
}

// GetBitN will read bit n as either true (1) or false (0).
// Error if n is not indexable in this BitSet.
func (bs *BitSet) GetBitN(n int) (bool, error) {
	byteOffset, bitOffset, offsetErr := bs.offsets(n)
	if offsetErr != nil {
		return false, offsetErr
	}
	return (0 != bs.bytesArr[byteOffset]&(1<<bitOffset)), nil
}
