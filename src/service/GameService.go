package service

import (
	"fmt"
	"time"
	"toy/src/logger"
)

type GameService struct {
	logger.LoggerTool
}

func (g *GameService) Start() {
	for i := 0; i < 10; i++ {
		fmt.Printf("game starting...%vs\r", i+1)
		time.Sleep(1e9)
	}

}
