package bootstrap

import (
	"GoAuth/src/api"
	"GoAuth/src/database"
	"GoAuth/src/hash"
	"GoAuth/src/models"
	"GoAuth/src/pkg/auth"
	"GoAuth/src/scheduler"
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

	// Initialize database
	err = database.Init()
	if err != nil {
		log.Fatalf("Database Service: Failed to Initialize. %v", err)
	}
	log.Println("Database Service: Initialized Successfully.")

	// Close Database
	defer func() {
		if err = database.GetInstance().Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
		log.Println("Database Service: Database Close Successfully.")
	}()

	// Migrate Models
	err = database.GetInstance().GetClient().AutoMigrate(models.Models()...)
	if err != nil {
		log.Fatalf("Database Service: Failed to Migrate. %v", err)
	}
	log.Println("Database Service: Database Migrate Successfully.")

	// Initialize HashPackage
	err = hash.GetInstance().Init()
	if err != nil {
		log.Fatalf("Hash Service: Failed to Initialize. %v", err)
	}
	log.Println("Hash Service: Initialized Successfully.")

	// Initialize AuthPackage
	err = auth.Init()
	if err != nil {
		log.Fatalf("Auth Service: Failed to Initialize. %v", err)
	}
	log.Println("Auth Service: Initialized Successfully.")

	// Initialize scheduler
	err = scheduler.Init()
	if err != nil {
		log.Fatalf("Scheduler Service: Failed to Initialize. %v", err)
	}
	log.Println("Scheduler Service: Initialized Successfully.")

	//Initialize API
	go func() {
		err = api.Init()
		if err != nil {
			log.Printf("API Service: Failed to Initialize. %v", err)
		}
	}()

	time.Sleep(50 * time.Millisecond)
	log.Println("Application is now running.\nPress CTRL-C to exit")
	<-sc

	log.Println("Application shutting down...")
	time.Sleep(1 * time.Second)
	return
}
