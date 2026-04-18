package main

import (
	"fmt"
	"net"
)

type NInt int

const (
	Sunday NInt = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(1 << 10)
}

func isUp(v net.Flags) bool {
	return v&net.FlagUp == net.FlagUp
}
