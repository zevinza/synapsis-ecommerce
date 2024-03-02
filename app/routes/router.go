package routes

import (
	"api/app/controller"
	"api/app/lib"
	"api/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"lab.tog.co.id/pub/pgsdk/config"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			lib.PrintStackTrace(e)
		},
	}))

	config.SetSecretKey("Edz0ThbXfQ6pYSQ3n267l1VQKGNbSuJE")

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Static("/swagger", "docs/swagger.json")
	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)
	api.Post("/logs", controller.PostLogs)

	// transaction
	transactionAPI := api.Group("/transactions")
	transactionAPI.Use(middleware.Oauth2Authentication)

}
