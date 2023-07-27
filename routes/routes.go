package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/phongpisut/fiber-air-htmx/models"
)

func PostHtmx(app *fiber.App, user *[]models.User) {

	app.Post("/", func(c *fiber.Ctx) error {

		name := c.FormValue("name")
		email := c.FormValue("email")
		tel := c.FormValue("tel")
		option := c.FormValue("option")
		convertOption, err := strconv.ParseUint(option, 10, 8)
		if err != nil {
			convertOption = 0
		}
		message := c.FormValue("message")

		newUser := models.User{
			Name:    name,
			Email:   email,
			Tel:     tel,
			Option:  convertOption,
			Message: message,
		}

		*user = append(*user, newUser)

		return c.Render("template/newUser", newUser)
	})
}
