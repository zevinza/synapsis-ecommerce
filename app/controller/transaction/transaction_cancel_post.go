package transaction

import (
	"api/app/lib"

	"github.com/gofiber/fiber/v2"
)

func PostTransactionCancel(c *fiber.Ctx) error {
	return lib.OK(c)
}
