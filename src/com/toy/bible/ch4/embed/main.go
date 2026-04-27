package main

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spoke = 20
	w = Wheel{
		Circle: Circle{Point{
			0,
			0,
		}, 1},
		Spoke: 0,
	}
}

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spoke int
}
