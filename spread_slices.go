package tools

import (
	"reflect"
)

// AnySlicer 不限类型的数组/切片 可提供方法有限
type AnySlicer[T any] []T

func ArrayAny[T any](arr []T) AnySlicer[T] {
	return arr
}

func ArrayAnys[T any](arr ...T) AnySlicer[T] {
	return arr
}

func (a AnySlicer[T]) Len() int {
	return len(a)
}

func (a AnySlicer[T]) Index(x T) int {
	for i, y := range a {
		if a.EqualItem(x, y) {
			return i
		}
	}
	return -1
}

func (a AnySlicer[T]) Exist(x T) bool {
	for _, y := range a {
		if a.EqualItem(x, y) {
			return true
		}
	}
	return false
}

func (a AnySlicer[T]) Equal(y AnySlicer[T]) bool {
	return reflect.DeepEqual(a, y)
}

func (a AnySlicer[T]) EqualItemIndex(index int, y T) bool {
	return reflect.DeepEqual(a[index], y)
}

func (a AnySlicer[T]) EqualItem(x, y T) bool {
	return reflect.DeepEqual(x, y)
}

func (a AnySlicer[T]) Merge(x AnySlicer[T]) AnySlicer[T] {
	return append(a, x...)
}

// Unique 数组去重, 如果数组元素的类型是 integer|float|string这些元素 请使用 ArrayOrdered.Unique 或 ArrayOrdered.UniqueOrdered
func (a AnySlicer[T]) Unique() AnySlicer[T] {
	var y = make(AnySlicer[T], 0)
	for _, x := range a {
		if !y.Exist(x) {
			y = append(y, x)
		}
	}
	return y
}
