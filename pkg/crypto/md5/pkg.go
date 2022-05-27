// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package md5

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("crypto/md5", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Funcs: map[string]func(c *internal.CallCtxt){
		"Sum": func(c *internal.CallCtxt) {

			data := c.Bytes(0)
			if c.Do() {
				c.Ret = Sum(data)
			}
		},
	},
	CUE: `{
	_
	exports: {
		Size?:      16
		BlockSize?: 64
		Sum: {
			in: [...#Arg] & [{
				name: "data"
				type: bytes | string
			}]
			out: bytes | string
		}
	}
}`,
}
