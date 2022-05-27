// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package hex

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("encoding/hex", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Funcs: map[string]func(c *internal.CallCtxt){
		"EncodedLen": func(c *internal.CallCtxt) {

			n := c.Int(0)
			if c.Do() {
				c.Ret = EncodedLen(n)
			}
		},
		"DecodedLen": func(c *internal.CallCtxt) {

			x := c.Int(0)
			if c.Do() {
				c.Ret = DecodedLen(x)
			}
		},
		"Decode": func(c *internal.CallCtxt) {

			s := c.String(0)
			if c.Do() {
				c.Ret, c.Err = Decode(s)
			}
		},
		"Dump": func(c *internal.CallCtxt) {

			data := c.Bytes(0)
			if c.Do() {
				c.Ret = Dump(data)
			}
		},
		"Encode": func(c *internal.CallCtxt) {

			src := c.Bytes(0)
			if c.Do() {
				c.Ret = Encode(src)
			}
		},
	},
	CUE: `{
	_
	exports: {
		Dump: {
			in: [...#Arg] & [{
				name: "data"
				type: bytes | string
			}]
			out: string
		}
		EncodedLen: {
			in: [...#Arg] & [{
				name: "n"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: >=-9223372036854775808 & <=9223372036854775807 & int
		}
		Encode: {
			in: [...#Arg] & [{
				name: "src"
				type: bytes | string
			}]
			out: string
		}
		DecodedLen: {
			in: [...#Arg] & [{
				name: "x"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: >=-9223372036854775808 & <=9223372036854775807 & int
		}
		Decode: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}]
			out: bytes | string
		}
	}
}`,
}
