package bitset 

import (
	"testing"
	"fmt"
)

func TestInit(t *testing.T) {
	bs := NewBitSet(8)
	if len(bs.bytes_arr) != 1 {
		t.Errorf("len of bytes should be 1 to hold 8 bits")
	}
	bs = NewBitSet(9)
	if len(bs.bytes_arr) != 2 {
		t.Errorf("len of bytes should be 2 to hold 9 bits")
	}
	bs = NewBitSet(17)
	if len(bs.bytes_arr) != 3 {
		t.Errorf("len of bytes should be 3 to hold 17 bits")
	}
	bs = NewBitSet(0)
	if len(bs.bytes_arr) != 0 {
		t.Errorf("len of bytes should be 0 to hold 0 bits")
	}
}

func TestSetUnsetGet(t *testing.T) {
	bs := NewBitSet(uint32(20))
 	set_err := bs.SetBitN(0)
 	if set_err != nil {
 		e := fmt.Sprintf("%s\n",set_err.Error())
		t.Errorf(e)
 	}
 	val,read_err := bs.GetBitN(0)
 	if read_err != nil {
 		e := fmt.Sprintf("%s\n",read_err.Error())
		t.Errorf(e)
 	}
	if val != true {
		t.Errorf("0 index should have read TRUE")
	}
	unset_err := bs.UnsetBitN(0)
 	if unset_err != nil {
 		e := fmt.Sprintf("%s\n",set_err.Error())
		t.Errorf(e)
 	}
 	val,read_err = bs.GetBitN(0)
 	if read_err != nil {
 		e := fmt.Sprintf("%s\n",read_err.Error())
		t.Errorf(e)
 	}
	if val != false {
		t.Errorf("0 index should have read FALSE")
	}
	set_err = bs.SetBitN(9)
 	if set_err != nil {
 		e := fmt.Sprintf("%s\n",set_err.Error())
		t.Errorf(e)
 	}
	set_err = bs.SetBitN(10)
 	if set_err != nil {
 		e := fmt.Sprintf("%s\n",set_err.Error())
		t.Errorf(e)
 	}
 	val,read_err = bs.GetBitN(9)
 	if read_err != nil {
 		e := fmt.Sprintf("%s\n",read_err.Error())
		t.Errorf(e)
 	}
	if val != true {
		t.Errorf("9 index should have read TRUE")
	}
	unset_err = bs.UnsetBitN(10)
 	if unset_err != nil {
 		e := fmt.Sprintf("%s\n",set_err.Error())
		t.Errorf(e)
 	}
 	val,read_err = bs.GetBitN(9)
 	if read_err != nil {
 		e := fmt.Sprintf("%s\n",read_err.Error())
		t.Errorf(e)
 	}
	if val != true {
		t.Errorf("9 index should have read TRUE")
	}
	unset_err = bs.UnsetBitN(9)
 	if unset_err != nil {
 		e := fmt.Sprintf("%s\n",set_err.Error())
		t.Errorf(e)
 	}
 	val,read_err = bs.GetBitN(9)
 	if read_err != nil {
 		e := fmt.Sprintf("%s\n",read_err.Error())
		t.Errorf(e)
 	}
	if val != false {
		t.Errorf("9 index should have read FALSE")
	}
}