package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, "Welcome to shutdown server graceful!")
	})
	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown", err)
	}
	log.Println("Server exiting...")
}
