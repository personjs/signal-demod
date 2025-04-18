package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/personjs/signal-demod/cmd/demod"
	"github.com/personjs/signal-demod/internal/config"
	"github.com/personjs/signal-demod/internal/services"
)

func main() {
	if err := config.Load(); err != nil {
		fmt.Println("failed to load config:", err)
		os.Exit(1)
	}

	services.InitLogger(config.App.Log)
	services.InitDatabase(config.App.DB)

	go StartFileServer()

	demod.Execute()
}

func StartFileServer() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	services.Logger.Info().Msg("Serving on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		services.Logger.Fatal().Err(err)
	}
}
