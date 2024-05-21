package main

import (
	"log/slog"
	"os"

	"github.com/1gkx/go-template-project/internal/app"
)

var build = "undefined"

func main() {

	if err := app.New(
		app.WithLogger(slog.Default()),
	).Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
