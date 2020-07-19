package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvins float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0.0
	BoilingC      Celsius = 100.0
	ZeroK         Kelvins = 0.0
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}

func (k Kelvins) String() string {
	return fmt.Sprintf("%gK", k)
}
