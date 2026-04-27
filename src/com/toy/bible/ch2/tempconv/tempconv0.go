package tempconv

import "fmt"

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度
const (
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC     Celsius = 0       //结冰点温度
	BiolingC      Celsius = 100     //沸水点温度
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
