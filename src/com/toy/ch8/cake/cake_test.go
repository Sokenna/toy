package cake_test

import (
	"github.com/Sokenna/toy/src/com/toy/ch8/cake"
	"testing"
	"time"
)

var defaults = cake.Shop{
	Cakes:        20,
	BakeTime:     time.Microsecond * 10,
	NumIcers:     1,
	IceTime:      time.Millisecond * 10,
	InscribeTime: time.Millisecond * 10,
}

func Benchmark(b *testing.B) {
	cakeshop := defaults
	cakeshop.Verbose = testing.Verbose()
	cakeshop.Work(b.N)
}

func BenchmarkBuffer(b *testing.B) {
	cakeshop := defaults
	cakeshop.Verbose = testing.Verbose()
	cakeshop.IceBuf = 10
	cakeshop.BakeBuf = 10
	cakeshop.Work(b.N)
}

func BenchmarkVariable(b *testing.B) {
	cakeshop := defaults
	cakeshop.Verbose = testing.Verbose()
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.Work(b.N) //252ms
}

func BenchmarkVariableBuffer(b *testing.B) {
	cakeshop := defaults
	cakeshop.Verbose = testing.Verbose()
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.IceBuf = 10
	cakeshop.BakeBuf = 10
	cakeshop.Work(b.N) //238ms
}
func BenchmarkSlowIcing(b *testing.B) {
	cakeshop := defaults
	cakeshop.Verbose = testing.Verbose()
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N) //1037ms
}
func BenchmarkSlowIcingManyIcers(b *testing.B) {
	cakeshop := defaults
	cakeshop.Verbose = testing.Verbose()
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N) //265ms
}
