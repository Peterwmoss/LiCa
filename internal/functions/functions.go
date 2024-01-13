package functions

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func ToHandler(component templ.Component) fiber.Handler {
  return adaptor.HTTPHandler(templ.Handler(component))
}

func IsHTMXRequest(ctx *fiber.Ctx) bool {
  if ctx.Get("HX-Request") != "" {
    return true
  }

  return false
}
