package main

import (
	"api/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	server := config.InitHttpHandler(fmt.Sprintf(":%s", os.Getenv("API_PORT")))

	go func() {
		log.Println("Server running on port:", os.Getenv("API_PORT"))
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf(err.Error())
			}
		}
	}()

	go func() {
		<-sigs
		// terminating, graceful shutdown
		log.Println("Shutting Down HTTP Server")
		_ = server.Shutdown(context.Background())
		done <- true
	}()

	<-done
	log.Println("Shutdown Gracefully")
}
