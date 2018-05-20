package main

import (
	"log"

	"github.com/futureperfect/todos/pkg/app"
	"github.com/futureperfect/todos/pkg/version"
)

func main() {
	app := app.NewApp()

	log.Printf("Starting todos server : %v", version.VERSION)

	app.Run()
}
