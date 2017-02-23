package bitset

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	bs := NewBitSet(8)
	if len(bs.bytesArr) != 1 {
		t.Errorf("len of bytes should be 1 to hold 8 bits")
	}
	bs = NewBitSet(9)
	if len(bs.bytesArr) != 2 {
		t.Errorf("len of bytes should be 2 to hold 9 bits")
	}
	bs = NewBitSet(17)
	if len(bs.bytesArr) != 3 {
		t.Errorf("len of bytes should be 3 to hold 17 bits")
	}
	bs = NewBitSet(0)
	if len(bs.bytesArr) != 0 {
		t.Errorf("len of bytes should be 0 to hold 0 bits")
	}
}

func TestSetUnsetGet(t *testing.T) {
	bs := NewBitSet(uint32(20))
	setErr := bs.SetBitN(0)
	if setErr != nil {
		e := fmt.Sprintf("%s\n", setErr.Error())
		t.Errorf(e)
	}
	val, readErr := bs.GetBitN(0)
	if readErr != nil {
		e := fmt.Sprintf("%s\n", readErr.Error())
		t.Errorf(e)
	}
	if val != true {
		t.Errorf("0 index should have read TRUE")
	}
	unsetErr := bs.UnsetBitN(0)
	if unsetErr != nil {
		e := fmt.Sprintf("%s\n", setErr.Error())
		t.Errorf(e)
	}
	val, readErr = bs.GetBitN(0)
	if readErr != nil {
		e := fmt.Sprintf("%s\n", readErr.Error())
		t.Errorf(e)
	}
	if val != false {
		t.Errorf("0 index should have read FALSE")
	}
	setErr = bs.SetBitN(9)
	if setErr != nil {
		e := fmt.Sprintf("%s\n", setErr.Error())
		t.Errorf(e)
	}
	setErr = bs.SetBitN(10)
	if setErr != nil {
		e := fmt.Sprintf("%s\n", setErr.Error())
		t.Errorf(e)
	}
	val, readErr = bs.GetBitN(9)
	if readErr != nil {
		e := fmt.Sprintf("%s\n", readErr.Error())
		t.Errorf(e)
	}
	if val != true {
		t.Errorf("9 index should have read TRUE")
	}
	unsetErr = bs.UnsetBitN(10)
	if unsetErr != nil {
		e := fmt.Sprintf("%s\n", setErr.Error())
		t.Errorf(e)
	}
	val, readErr = bs.GetBitN(9)
	if readErr != nil {
		e := fmt.Sprintf("%s\n", readErr.Error())
		t.Errorf(e)
	}
	if val != true {
		t.Errorf("9 index should have read TRUE")
	}
	unsetErr = bs.UnsetBitN(9)
	if unsetErr != nil {
		e := fmt.Sprintf("%s\n", setErr.Error())
		t.Errorf(e)
	}
	val, readErr = bs.GetBitN(9)
	if readErr != nil {
		e := fmt.Sprintf("%s\n", readErr.Error())
		t.Errorf(e)
	}
	if val != false {
		t.Errorf("9 index should have read FALSE")
	}
}
