package lenconv

import "fmt"

type Foot float64
type Meter float64

const (
	OneMeter Meter = 1
	OneFoot  Meter = 0.3048
)

func (l Foot) String() string {
	return fmt.Sprintf("%gft", l)
}
func (l Meter) String() string {
	return fmt.Sprintf("%gft", l)
}
