package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/munsifali/student-api/internal/config"
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
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student Api"))
	})

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

	slog.Info("Shutting Down Server!")

}
