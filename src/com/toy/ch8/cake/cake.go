package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool
	Cakes          int           // 每批次烤多少个蛋糕
	BakeTime       time.Duration // 烤蛋糕的标准时间
	BakeStdDev     time.Duration // 烤蛋糕的标准偏差
	BakeBuf        int           // 每批次烤蛋糕的缓冲数
	NumIcers       int           // 冰激凌分发者的数量
	IceTime        time.Duration // 给一块蛋糕裱花的时间
	IceStdDev      time.Duration // 给一块蛋糕裱花的时间偏差
	IceBuf         int           // 裱花与刻字工序间的缓冲槽位
	InscribeTime   time.Duration // 给一块蛋糕刻字的时间
	InscribeStdDev time.Duration // 给一块蛋糕刻字的时间偏差
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscriber", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			s.icer(iced, baked)
		}
	}
}

func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
