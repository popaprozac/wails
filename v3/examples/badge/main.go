package main

import (
	_ "embed"
	"log"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

func main() {

	app := application.New(application.Options{
		Name:        "badging",
		Description: "A demo of using raw HTML & CSS",
	})

	app.OnApplicationEvent(events.Common.ApplicationStarted, func(event *application.ApplicationEvent) {
		time.Sleep(time.Second * 1)
		app.SetBadge("1")
		time.Sleep(time.Second * 5)
		app.SetBadge("label")
		time.Sleep(time.Second * 5)
		app.SetBadge("")
	})

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
