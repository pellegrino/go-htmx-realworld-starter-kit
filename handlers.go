package main

import "github.com/gofiber/fiber/v2"

// indexViewHandler handles a view for the index page.
func indexViewHandler(c *fiber.Ctx) error {
	return c.Render("pages/index", nil)
}

// showContentAPIHandler handles an API endpoint to show content.
func showContentAPIHandler(c *fiber.Ctx) error {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if c.Get("HX-Request") != "true" {
		// If not, return HTTP 400 error.
		return fiber.NewError(fiber.StatusBadRequest, "non-htmx request")
	}

	return c.SendString("<p>🎉 Hello from <strong>htmx</strong>!<br>(<code>GET /api/show</code>)</p>")
}
