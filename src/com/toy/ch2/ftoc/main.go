package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g° = %g°C\n", freezingF, fToC(freezingF)) //"32°F = 0°C"
	fmt.Printf("%g° = %g°C\n", boilingF, fToC(boilingF))   //"212°F = 100°C"

}
func fToC(c float64) float64 {
	return (c - 32) * 5 / 9
}
