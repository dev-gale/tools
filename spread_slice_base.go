package tools

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"strings"
)

type Slicer[T constraints.Ordered] []T

func Array[T constraints.Ordered](a []T) Slicer[T] {
	return a
}
func Arrays[T constraints.Ordered](a ...T) Slicer[T] {
	return a
}

func (a Slicer[T]) Len() int {
	return len(a)
}

func (a Slicer[T]) Index(x T) int {
	for i, y := range a {
		if x == y {
			return i
		}
	}
	return -1
}

func (a Slicer[T]) Exist(x T) bool {
	for _, y := range a {
		if x == y {
			return true
		}
	}
	return false
}

func (a Slicer[T]) Equal(y Slicer[T]) bool {
	l := a.Len()
	if l != y.Len() {
		return false
	}

	for i := 0; i < l; i++ {
		if a[i] != y[i] {
			return false
		}
	}
	return true
}

func (a Slicer[T]) Merge(x Slicer[T]) Slicer[T] {
	return append(a, x...)
}

// UniqueFast 不保留原数组顺序
func (a Slicer[T]) UniqueFast() Slicer[T] {
	m := make(map[T]bool)
	for _, x := range a {
		if _, ok := m[x]; !ok {
			m[x] = true
		}
	}
	var y Slicer[T]
	for k := range m {
		y = append(y, k)
	}
	return y
}

// UniqueSort 保留原数组顺序
func (a Slicer[T]) UniqueSort() Slicer[T] {
	var y = make(Slicer[T], 0)
	for _, x := range a {
		if !y.Exist(x) {
			y = append(y, x)
		}
	}
	return y
}

func (a Slicer[T]) Remove(index int) Slicer[T] {
	return slices.Delete(a, index, index+1)
}

func (a Slicer[T]) RemoveValue(value T) Slicer[T] {
	return slices.DeleteFunc(a, func(v T) bool {
		return v == value
	})
}

func (a Slicer[T]) Push(v T) Slicer[T] {
	return append(a, v)
}

func (a Slicer[T]) Replace(oldVal, newVal T) Slicer[T] {
	if a.Len() > 0 {
		for i, x := range a {
			if x == oldVal {
				a[i] = newVal
			}
		}
	}
	return a
}

func (a Slicer[T]) Join(sep string) string {
	if a.Len() == 0 {
		return ""
	}
	return strings.Join(a.ToString(), sep)
}

func (a Slicer[T]) ToString() Slicer[string] {
	var y = make([]string, a.Len())
	for i, v := range a {
		y[i] = fmt.Sprint(v)
	}
	return y
}

func (a Slicer[T]) ToInt() Slicer[int] {
	var x any = a
	switch x.(type) {
	case Slicer[int]:
		return x.(Slicer[int])
	case Slicer[int8]:
		return ArrayInteger2Int(x.(Slicer[int8]))
	case Slicer[int16]:
		return ArrayInteger2Int(x.(Slicer[int16]))
	case Slicer[int32]:
		return ArrayInteger2Int(x.(Slicer[int32]))
	case Slicer[int64]:
		return ArrayInteger2Int(x.(Slicer[int64]))
	case Slicer[uint]:
		return ArrayInteger2Int(x.(Slicer[uint]))
	case Slicer[uint8]:
		return ArrayInteger2Int(x.(Slicer[uint8]))
	case Slicer[uint16]:
		return ArrayInteger2Int(x.(Slicer[uint16]))
	case Slicer[uint32]:
		return ArrayInteger2Int(x.(Slicer[uint32]))
	case Slicer[uint64]:
		return ArrayInteger2Int(x.(Slicer[uint64]))
	case Slicer[float32]:
		return ArrayFloat2Int(x.(Slicer[float32]))
	case Slicer[float64]:
		return ArrayFloat2Int(x.(Slicer[float64]))
	case Slicer[string]:
		return ArrayString2Int(x.(Slicer[string]))
	default:
		return nil
	}
}

func (a Slicer[T]) ToInt64() Slicer[int64] {
	var x any = a
	switch x.(type) {
	case Slicer[int]:
		return ArrayInteger2Int64(x.(Slicer[int]))
	case Slicer[int8]:
		return ArrayInteger2Int64(x.(Slicer[int8]))
	case Slicer[int16]:
		return ArrayInteger2Int64(x.(Slicer[int16]))
	case Slicer[int32]:
		return ArrayInteger2Int64(x.(Slicer[int32]))
	case Slicer[int64]:
		return x.(Slicer[int64])
	case Slicer[uint]:
		return ArrayInteger2Int64(x.(Slicer[uint]))
	case Slicer[uint8]:
		return ArrayInteger2Int64(x.(Slicer[uint8]))
	case Slicer[uint16]:
		return ArrayInteger2Int64(x.(Slicer[uint16]))
	case Slicer[uint32]:
		return ArrayInteger2Int64(x.(Slicer[uint32]))
	case Slicer[uint64]:
		return ArrayInteger2Int64(x.(Slicer[uint64]))
	case Slicer[float32]:
		return ArrayFloat2Int64(x.(Slicer[float32]))
	case Slicer[float64]:
		return ArrayFloat2Int64(x.(Slicer[float64]))
	case Slicer[string]:
		return ArrayString2Int64(x.(Slicer[string]))
	default:
		return nil
	}
}

func (a Slicer[T]) ToFloat64() Slicer[float64] {
	var x any = a
	switch x.(type) {
	case Slicer[int]:
		return ArrayInteger2Float64(x.(Slicer[int]))
	case Slicer[int8]:
		return ArrayInteger2Float64(x.(Slicer[int8]))
	case Slicer[int16]:
		return ArrayInteger2Float64(x.(Slicer[int16]))
	case Slicer[int32]:
		return ArrayInteger2Float64(x.(Slicer[int32]))
	case Slicer[int64]:
		return ArrayInteger2Float64(x.(Slicer[int64]))
	case Slicer[uint]:
		return ArrayInteger2Float64(x.(Slicer[uint]))
	case Slicer[uint8]:
		return ArrayInteger2Float64(x.(Slicer[uint8]))
	case Slicer[uint16]:
		return ArrayInteger2Float64(x.(Slicer[uint16]))
	case Slicer[uint32]:
		return ArrayInteger2Float64(x.(Slicer[uint32]))
	case Slicer[uint64]:
		return ArrayInteger2Float64(x.(Slicer[uint64]))
	case Slicer[float32]:
		return ArrayFloat2Float64(x.(Slicer[float32]))
	case Slicer[float64]:
		return x.(Slicer[float64])
	case Slicer[string]:
		return ArrayString2Float64(x.(Slicer[string]))
	default:
		return nil
	}
}
