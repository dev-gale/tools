package tools

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

type Integer int64

func Int[T constraints.Integer](v T) Integer {
	return Integer(v)
}

func (i Integer) Int() int {
	return int(i)
}

func (i Integer) Int64() int64 {
	return int64(i)
}

func (i Integer) Uint() uint {
	return uint(i)
}

func (i Integer) Uint64() uint64 {
	return uint64(i)
}

func (i Integer) Floater() Floater {
	return Floater(i)
}

func (i Integer) Stringer() Stringer {
	return Stringer(strconv.Itoa(i.Int()))
}

func (i Integer) IsEmpty() bool {
	return i.Int() == 0
}
