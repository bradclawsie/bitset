package bitset

import (
	"fmt"
	"errors"
	"strconv"
	"strings"
	"sync"
)

// According to the Go spec, byte = uint8, so we set our byte width to 8 here
const (
	WIDTH = 8
)

// The bitset is simply an array of bytes large enough to contain the number of bits the
// user requests. Values are encoded in a little endian fashion: in a single bytes, bits 
// would be indexes as 7 | 6 | 5 | 4 | 3 | 2 | 1 | 0. Callers need not concern themselves
// with this implementation detail: getting, setting and unsetting bits is done through
// functions.
type BitSet struct {
	bytes_arr []byte
	// lock for access
	mutex sync.RWMutex
	size int
}

// Create a new BitSet that can accomodate n bits.
func NewBitSet(n uint32) *BitSet {
	bs := new(BitSet)
	if n == 0 {
		bs.bytes_arr = make([]byte,0)
	} else {
		// round up so that the number of bytes can accomodate n bits
		bs.bytes_arr = make([]byte,((n - 1) / WIDTH) + 1)
	}
	bs.size = len(bs.bytes_arr) * WIDTH
	return bs
}

// return the size of the bitset
func (bs *BitSet) Size() int {
	return bs.size
}

// Create a viewable form of the BitSet
func (bs *BitSet) DumpBitSet() string {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()
	pss := ""
	for _,v := range bs.bytes_arr {
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

// Calculate byte and bit offsets for operations. Fails if n is not indexable in this BitSet.
func (bs *BitSet) offsets(n int) (int,byte,error) {
	if n < 0 || n > (len(bs.bytes_arr) * WIDTH) {
		e := fmt.Sprintf("index %d under or overflows bitset",n)
		return 0,0,errors.New(e)
	}
	byte_offset := len(bs.bytes_arr)-(n/WIDTH)-1
	bit_offset := byte(n % WIDTH)
	return byte_offset,bit_offset,nil
}

// Set bit n to be 1. Error if n is not indexable in this BitSet.
func (bs *BitSet) SetBitN(n int) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()
	byte_offset,bit_offset,offset_err := bs.offsets(n)
	if offset_err != nil {
		return offset_err
	}
	bs.bytes_arr[byte_offset] |= (1 << bit_offset) 
	return nil
}

// Set bit n to be 0. Error is n is not indexable in this BitSet.
func (bs *BitSet) UnsetBitN(n int) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()
	byte_offset,bit_offset,offset_err := bs.offsets(n)
	if offset_err != nil {
		return offset_err
	}
	bs.bytes_arr[byte_offset] ^= (1 << bit_offset) 
	return nil
}

// Read bit n as either true (1) or false (0). Error is n is not indexable in this BitSet.
func (bs *BitSet) GetBitN(n int) (bool,error) {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()
	byte_offset,bit_offset,offset_err := bs.offsets(n)
	if offset_err != nil {
	 	return false,offset_err
	}
	return (0 != bs.bytes_arr[byte_offset] & (1 << bit_offset)),nil
}
