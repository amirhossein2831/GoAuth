package bootstrap

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init() (err error) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	log.Println("Application is starting...")

	

	log.Println("Application is now running.\nPress CTRL-C to exit")
	<-sc

	log.Println("Application shutting down...")
	time.Sleep(1 * time.Second)
	return
}
