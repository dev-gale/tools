package tools

import "github.com/shopspring/decimal"

type Decimaler struct {
	f decimal.Decimal
}

func Decimal(f float64) Decimaler {
	return Decimaler{f: decimal.NewFromFloat(f)}
}

func (x Decimaler) String() string {
	return x.f.String()
}

func (x Decimaler) Add(y float64) Decimaler {
	x.f = x.f.Add(decimal.NewFromFloat(y))
	return x
}

func (x Decimaler) Sub(y float64) Decimaler {
	x.f = x.f.Sub(decimal.NewFromFloat(y))
	return x
}

func (x Decimaler) Mul(y float64) Decimaler {
	x.f = x.f.Mul(decimal.NewFromFloat(y))
	return x
}

func (x Decimaler) Div(y float64) Decimaler {
	x.f = x.f.Div(decimal.NewFromFloat(y))
	return x
}

func (x Decimaler) Pow(y float64) Decimaler {
	x.f = x.f.Pow(decimal.NewFromFloat(y))
	return x
}

func (x Decimaler) Float() float64 {
	f, _ := x.f.Float64()
	return f
}

func (x Decimaler) FloatX() Floater {
	return Floater(x.Float())
}

func (x Decimaler) IsZero() bool {
	return x.f.IsZero()
}
