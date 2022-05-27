// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package strings

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("strings", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Funcs: map[string]func(c *internal.CallCtxt){
		"ByteAt": func(c *internal.CallCtxt) {

			b, i := c.Bytes(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = ByteAt(b, i)
			}
		},
		"ByteSlice": func(c *internal.CallCtxt) {

			b, start, end := c.Bytes(0), c.Int(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = ByteSlice(b, start, end)
			}
		},
		"Runes": func(c *internal.CallCtxt) {

			s := c.String(0)
			if c.Do() {
				c.Ret = Runes(s)
			}
		},
		"MinRunes": func(c *internal.CallCtxt) {

			s, min := c.String(0), c.Int(1)
			if c.Do() {
				c.Ret = MinRunes(s, min)
			}
		},
		"MaxRunes": func(c *internal.CallCtxt) {

			s, max := c.String(0), c.Int(1)
			if c.Do() {
				c.Ret = MaxRunes(s, max)
			}
		},
		"ToTitle": func(c *internal.CallCtxt) {

			s := c.String(0)
			if c.Do() {
				c.Ret = ToTitle(s)
			}
		},
		"ToCamel": func(c *internal.CallCtxt) {

			s := c.String(0)
			if c.Do() {
				c.Ret = ToCamel(s)
			}
		},
		"SliceRunes": func(c *internal.CallCtxt) {

			s, start, end := c.String(0), c.Int(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = SliceRunes(s, start, end)
			}
		},
		"Compare": func(c *internal.CallCtxt) {

			a, b := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = Compare(a, b)
			}
		},
		"Count": func(c *internal.CallCtxt) {

			s, substr := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = Count(s, substr)
			}
		},
		"Contains": func(c *internal.CallCtxt) {

			s, substr := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = Contains(s, substr)
			}
		},
		"ContainsAny": func(c *internal.CallCtxt) {

			s, chars := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = ContainsAny(s, chars)
			}
		},
		"LastIndex": func(c *internal.CallCtxt) {

			s, substr := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = LastIndex(s, substr)
			}
		},
		"IndexAny": func(c *internal.CallCtxt) {

			s, chars := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = IndexAny(s, chars)
			}
		},
		"LastIndexAny": func(c *internal.CallCtxt) {

			s, chars := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = LastIndexAny(s, chars)
			}
		},
		"SplitN": func(c *internal.CallCtxt) {

			s, sep, n := c.String(0), c.String(1), c.Int(2)
			if c.Do() {
				c.Ret = SplitN(s, sep, n)
			}
		},
		"SplitAfterN": func(c *internal.CallCtxt) {

			s, sep, n := c.String(0), c.String(1), c.Int(2)
			if c.Do() {
				c.Ret = SplitAfterN(s, sep, n)
			}
		},
		"Split": func(c *internal.CallCtxt) {

			s, sep := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = Split(s, sep)
			}
		},
		"SplitAfter": func(c *internal.CallCtxt) {

			s, sep := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = SplitAfter(s, sep)
			}
		},
		"Fields": func(c *internal.CallCtxt) {

			s := c.String(0)
			if c.Do() {
				c.Ret = Fields(s)
			}
		},
		"Join": func(c *internal.CallCtxt) {

			elems, sep := c.StringList(0), c.String(1)
			if c.Do() {
				c.Ret = Join(elems, sep)
			}
		},
		"HasPrefix": func(c *internal.CallCtxt) {

			s, prefix := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = HasPrefix(s, prefix)
			}
		},
		"HasSuffix": func(c *internal.CallCtxt) {

			s, suffix := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = HasSuffix(s, suffix)
			}
		},
		"Repeat": func(c *internal.CallCtxt) {

			s, count := c.String(0), c.Int(1)
			if c.Do() {
				c.Ret = Repeat(s, count)
			}
		},
		"ToUpper": func(c *internal.CallCtxt) {

			s := c.String(0)
			if c.Do() {
				c.Ret = ToUpper(s)
			}
		},
		"ToLower": func(c *internal.CallCtxt) {

			s := c.String(0)
			if c.Do() {
				c.Ret = ToLower(s)
			}
		},
		"Trim": func(c *internal.CallCtxt) {

			s, cutset := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = Trim(s, cutset)
			}
		},
		"TrimLeft": func(c *internal.CallCtxt) {

			s, cutset := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = TrimLeft(s, cutset)
			}
		},
		"TrimRight": func(c *internal.CallCtxt) {

			s, cutset := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = TrimRight(s, cutset)
			}
		},
		"TrimSpace": func(c *internal.CallCtxt) {

			s := c.String(0)
			if c.Do() {
				c.Ret = TrimSpace(s)
			}
		},
		"TrimPrefix": func(c *internal.CallCtxt) {

			s, prefix := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = TrimPrefix(s, prefix)
			}
		},
		"TrimSuffix": func(c *internal.CallCtxt) {

			s, suffix := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = TrimSuffix(s, suffix)
			}
		},
		"Replace": func(c *internal.CallCtxt) {

			s, old, new, n := c.String(0), c.String(1), c.String(2), c.Int(3)
			if c.Do() {
				c.Ret = Replace(s, old, new, n)
			}
		},
		"Index": func(c *internal.CallCtxt) {

			s, substr := c.String(0), c.String(1)
			if c.Do() {
				c.Ret = Index(s, substr)
			}
		},
	},
	CUE: `{
	_
	exports: {
		TrimSuffix: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "suffix"
				type: string
			}]
			out: string
		}
		TrimSpace: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}]
			out: string
		}
		TrimRight: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "cutset"
				type: string
			}]
			out: string
		}
		TrimPrefix: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "prefix"
				type: string
			}]
			out: string
		}
		TrimLeft: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "cutset"
				type: string
			}]
			out: string
		}
		Trim: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "cutset"
				type: string
			}]
			out: string
		}
		ToUpper: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}]
			out: string
		}
		ToTitle: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}]
			out: string
		}
		ToLower: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}]
			out: string
		}
		ToCamel: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}]
			out: string
		}
		SplitN: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "sep"
				type: string
			}, {
				name: "n"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: [...string]
		}
		SplitAfterN: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "sep"
				type: string
			}, {
				name: "n"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: [...string]
		}
		SplitAfter: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "sep"
				type: string
			}]
			out: [...string]
		}
		SliceRunes: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "start"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}, {
				name: "end"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: string
		}
		Runes: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}]
			out: [...]
		}
		Replace: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "old"
				type: string
			}, {
				name: "new"
				type: string
			}, {
				name: "n"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: string
		}
		MinRunes: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "min"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: bool
		}
		MaxRunes: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "max"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: bool
		}
		LastIndexAny: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "chars"
				type: string
			}]
			out: >=-9223372036854775808 & <=9223372036854775807 & int
		}
		LastIndex: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "substr"
				type: string
			}]
			out: >=-9223372036854775808 & <=9223372036854775807 & int
		}
		Join: {
			in: [...#Arg] & [{
				name: "elems"
				type: [...string]
			}, {
				name: "sep"
				type: string
			}]
			out: string
		}
		IndexAny: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "chars"
				type: string
			}]
			out: >=-9223372036854775808 & <=9223372036854775807 & int
		}
		Index: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "substr"
				type: string
			}]
			out: >=-9223372036854775808 & <=9223372036854775807 & int
		}
		HasSuffix: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "suffix"
				type: string
			}]
			out: bool
		}
		HasPrefix: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "prefix"
				type: string
			}]
			out: bool
		}
		Fields: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}]
			out: [...string]
		}
		Count: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "substr"
				type: string
			}]
			out: >=-9223372036854775808 & <=9223372036854775807 & int
		}
		ContainsAny: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "chars"
				type: string
			}]
			out: bool
		}
		Compare: {
			in: [...#Arg] & [{
				name: "a"
				type: string
			}, {
				name: "b"
				type: string
			}]
			out: >=-9223372036854775808 & <=9223372036854775807 & int
		}
		ByteSlice: {
			in: [...#Arg] & [{
				name: "b"
				type: bytes | string
			}, {
				name: "start"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}, {
				name: "end"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: bytes | string
		}
		ByteAt: {
			in: [...#Arg] & [{
				name: "b"
				type: bytes | string
			}, {
				name: "i"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: >=0 & <=255 & int
		}
		Repeat: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "count"
				type: >=-9223372036854775808 & <=9223372036854775807 & int
			}]
			out: string
		}
		Contains: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "substr"
				type: string
			}]
			out: bool
		}
		Split: {
			in: [...#Arg] & [{
				name: "s"
				type: string
			}, {
				name: "sep"
				type: string
			}]
			out: [...string]
		}
	}
}`,
}
