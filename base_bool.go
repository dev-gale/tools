package tools

type Boolean bool

func Bool(v bool) Boolean {
	b := Boolean(v)
	return b
}

func (b Boolean) Bool() bool {
	return bool(b)
}

func (b Boolean) IsEmpty() bool {
	return !b.Bool()
}

func (b Boolean) Integer() Integer {
	return Int(Ternary(b.Bool(), 1, 0))
}

func (b Boolean) Floater() Floater {
	return Float(Ternary(b.Bool(), 1, 0.0))
}

func (b Boolean) Stringer() Stringer {
	return Stringer(Ternary(b.Bool(), "true", "false"))
}

func (b Boolean) Int() int {
	return b.Integer().Int()
}

func (b Boolean) Int64() int64 {
	return b.Integer().Int64()
}

func (b Boolean) Float64() float64 {
	return b.Floater().Float64()
}

func (b Boolean) String() string {
	return b.Stringer().String()
}

func (b Boolean) Pointer() *bool {
	p := b.Bool()
	return &p
}
