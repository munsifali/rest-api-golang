package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/munsifali/student-api/internal/config"
	"github.com/munsifali/student-api/internal/config/http/handlers/student"
)

func main() {
	/*
		Load Configuration
		Setup Database
		Setup Router
		Setup Http Server
	*/
	configuration := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("GET /api/student", student.CreateStudent())

	server := http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Faild to start server!")
		}
	}()

	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Faild to Shut Down server!")
	}
	slog.Info("Shutting Down Server!")
}
