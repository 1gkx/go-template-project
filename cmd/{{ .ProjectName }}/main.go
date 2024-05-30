package main

import (
	"log/slog"
	"os"

	"{{ .ModeleName }}/internal/app"
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
