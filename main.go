package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/phongpisut/fiber-air-htmx/models"
	"github.com/phongpisut/fiber-air-htmx/routes"
)

func main() {
	engine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./public/static")

	app.Get("/api", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg)
	})

	user := []models.User{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Users": user,
		})
	})

	routes.PostHtmx(app, &user)

	app.Listen(":3000")

}
