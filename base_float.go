package tools

import (
	"fmt"
	"math"
)

type Floater float64

func Float[T float32 | float64](v T) Floater {
	return Floater(v)
}

func (v Floater) Float64() float64 {
	return float64(v)
}

func (v Floater) Ceil(n float64) Floater {
	pow := math.Pow(10, n)
	return Floater(math.Ceil(float64(v)*pow) / pow)
}

func (v Floater) Floor(n float64) Floater {
	pow := math.Pow(10, n)
	return Floater(math.Floor(float64(v)*pow) / pow)
}

func (v Floater) Round(n float64) Floater {
	pow := math.Pow(10, n)
	return Floater(math.Round(float64(v)*pow) / pow)
}

func (v Floater) Format(n int) string {
	format := fmt.Sprintf("%%.%df", n)
	return fmt.Sprintf(format, v)
}

func (v Floater) Integer() Integer {
	return Integer(v)
}

func (v Floater) Stringer() Stringer {
	return Stringer(fmt.Sprintf("%f", v.Float64()))
}

func (v Floater) Pointer() *float64 {
	p := v.Float64()
	return &p
}

func (v Floater) IsEmpty() bool {
	return v.Float64() == 0
}
