package tools

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
)

func ArrayColumn[T, V any](array []T, k string) []V {
	l := len(array)
	if l == 0 {
		return nil
	}
	values := make([]V, len(array))
	switch reflect.TypeOf(array).Elem().Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < l; i++ {
			values[i] = reflect.ValueOf(array[i]).Index(int(reflect.ValueOf(k).Int())).Interface().(V)
		}
		break
	case reflect.Map:
		for i := 0; i < l; i++ {
			values[i] = reflect.ValueOf(array[i]).MapIndex(reflect.ValueOf(k)).Interface().(V)
		}
		break
	case reflect.Struct:
		for i := 0; i < l; i++ {
			values[i] = reflect.ValueOf(array[i]).FieldByName(reflect.ValueOf(k).String()).Interface().(V)
		}
		break
	default:
		return nil
	}
	return values
}

type Number interface {
	constraints.Integer | constraints.Float
}

func ArrayInteger2Int64[T constraints.Integer](in Slicer[T]) Slicer[int64] {
	out := make(Slicer[int64], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = int64(in[i])
	}
	return out
}

func ArrayFloat2Int64[T constraints.Float](in Slicer[T]) Slicer[int64] {
	out := make(Slicer[int64], in.Len())
	for i := 0; i < in.Len(); i++ {
		v, _ := strconv.Atoi(fmt.Sprintf("%.0f", in[i]))
		out[i] = int64(v)
	}
	return out
}

func ArrayString2Int64(in Slicer[string]) Slicer[int64] {
	out := make(Slicer[int64], in.Len())
	for i := 0; i < in.Len(); i++ {
		v, _ := strconv.Atoi(in[i])
		out[i] = int64(v)
	}
	return out
}

func ArrayInteger2Int[T constraints.Integer](in Slicer[T]) Slicer[int] {
	out := make(Slicer[int], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = int(in[i])
	}
	return out
}

func ArrayFloat2Int[T constraints.Float](in Slicer[T]) Slicer[int] {
	out := make(Slicer[int], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i], _ = strconv.Atoi(fmt.Sprintf("%.0f", in[i]))
	}
	return out
}

func ArrayString2Int(in Slicer[string]) Slicer[int] {
	out := make(Slicer[int], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i], _ = strconv.Atoi(in[i])
	}
	return out
}

func ArrayNumber2String[T Number](in Slicer[T]) Slicer[string] {
	out := make(Slicer[string], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = fmt.Sprint(in[i])
	}
	return out
}

func ArrayInteger2Float64[T constraints.Integer](in Slicer[T]) Slicer[float64] {
	out := make(Slicer[float64], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = float64(in[i])
	}
	return out
}

func ArrayFloat2Float64[T constraints.Float](in Slicer[T]) Slicer[float64] {
	out := make(Slicer[float64], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = float64(in[i])
	}
	return out
}

func ArrayString2Float64(in Slicer[string]) Slicer[float64] {
	out := make(Slicer[float64], in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i], _ = strconv.ParseFloat(in[i], 64)
	}
	return out
}
