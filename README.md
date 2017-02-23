[![License BSD](https://img.shields.io/badge/License-BSD-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)
[![Go Report Card](https://goreportcard.com/badge/github.com/bradclawsie/bitset)](https://goreportcard.com/report/github.com/bradclawsie/bitset)
[![GoDoc](https://godoc.org/github.com/bradclawsie/httpshutdown?status.svg)](http://godoc.org/github.com/bradclawsie/bitset)
[![Build Status](https://travis-ci.org/bradclawsie/bitset.png)](https://travis-ci.org/bradclawsie/bitset)

## bitset

This package implements a bitset in Go. By a bitset, I mean a structure that encodes
true/false values as 1/0 inside bytes. Eight boolean values should be accomodated by one byte 
using this structure.

To allow for large bitsets, the constructor asks for a number of bits you
wish to contain, and then constructs a []byte with size suitable to hold this number. 

Values are encoded in a little endian fashion: in a single bytes, bits 
would be indexes as 7 | 6 | 5 | 4 | 3 | 2 | 1 | 0. Callers need not concern themselves
with this implementation detail: getting, setting and unsetting bits is done through
functions.

There are three main functions: 

`SetBitN(n)` which sets the bit at position `n` to `1` (`true`)

`UnsetBitN(n)` which sets the bit at position `n` to `0` (`false`)

`GetBitN(n)` which gets the bit as position `n` and returns it as a bool (`1` = `true`, `0` = `false`)

For those wishing to visualize the layout of bits, a convenience function called DumpBitSet
is provided.


