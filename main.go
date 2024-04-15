/* Main (Entrypoint)

Connects to the database
Sets up the routes in a mux
Starts the server in a goroutine
Gracefully shuts down the server on termination signal

Joaquin Badillo
2024-04-15
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	manager "github.com/JoaquinBadillo/FlightsAPI/db/provider"
	flights "github.com/JoaquinBadillo/FlightsAPI/routes"
)

func Colorize(color int, message string) string {
	return fmt.Sprintf("\033[0;%dm%s\033[0m", color, message)
}

func main() {
	manager.Connect()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/flights", flights.GetFlights)
	mux.HandleFunc("GET /api/flights/{id}", flights.GetFlight)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("⚡ Server running")
		if os.Getenv("PRODUCTION") == "" {
			log.Println(Colorize(34, fmt.Sprintf("   http://localhost%s", port)))
		}
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-sigChan
	log.Println("🕊️ Gracefully shutting down")
	manager.Mgr.Close()
	server.Shutdown(context.Background())
}
