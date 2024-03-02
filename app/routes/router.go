package routes

import (
	"api/app/controller"
	"api/app/controller/cart"
	"api/app/controller/category"
	"api/app/controller/product"
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

	// Cart
	cartAPI := api.Group("/carts")
	cartAPI.Use(middleware.TokenValidator())
	cartAPI.Post("/", cart.PostCart)
	cartAPI.Get("/", cart.GetCart)
	cartAPI.Put("/:id", cart.PutCart)
	cartAPI.Get("/:id", cart.GetCartID)
	cartAPI.Delete("/:id", cart.DeleteCart)

	// Category
	categoryAPI := api.Group("/categories")
	categoryAPI.Use(middleware.TokenValidator())
	categoryAPI.Post("/", category.PostCategory)
	categoryAPI.Get("/", category.GetCategory)
	categoryAPI.Put("/:id", category.PutCategory)
	categoryAPI.Get("/:id", category.GetCategoryID)
	categoryAPI.Delete("/:id", category.DeleteCategory)

	// Product
	productAPI := api.Group("/products")
	productAPI.Use(middleware.TokenValidator())
	productAPI.Post("/", product.PostProduct)
	productAPI.Get("/", product.GetProduct)
	productAPI.Put("/:id", product.PutProduct)
	productAPI.Get("/:id", product.GetProductID)
	productAPI.Delete("/:id", product.DeleteProduct)

	// transaction
	transactionAPI := api.Group("/transactions")
	transactionAPI.Use(middleware.Oauth2Authentication)

}
