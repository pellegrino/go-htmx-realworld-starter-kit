package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables.
	port := 5050
	readTimeout := 5
	writeTimeout := 10

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	config := fiber.Config{
		Views:        html.NewFileSystem(http.Dir("./templates"), ".html"),
		ViewsLayout:  "main",
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	// Create a new Fiber server.
	server := fiber.New(config)

	// Add Fiber middlewares.
	server.Use(logger.New())

	// Handle static files.
	server.Static("/static", "./static")

	// Handle index page view.
	server.Get("/", indexViewHandler)

	// Handle API endpoints.
	server.Get("/api/show", showContentAPIHandler)

	return server.Listen(fmt.Sprintf(":%d", port))
}
