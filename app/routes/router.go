package routes

import (
	"api/app/controller"
	"api/app/controller/account"
	"api/app/controller/cart"
	"api/app/controller/category"
	"api/app/controller/product"
	"api/app/controller/transaction"
	"api/app/controller/user"
	"api/app/lib"
	"api/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
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

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Static("/swagger", "docs/swagger.json")
	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)
	api.Post("/logs", controller.PostLogs)

	// Account
	accountAPI := api.Group("/accounts")
	accountAPI.Use(middleware.TokenValidator())
	accountAPI.Post("/register", account.PostAccountRegister)
	accountAPI.Post("/login", account.PostAccountLogin)
	accountAPI.Post("/logout", account.PostAccountLogout)
	accountAPI.Post("/refresh", account.PostAccountRefresh)

	// Cart
	cartAPI := api.Group("/carts")
	cartAPI.Use(middleware.Oauth2Authentication)
	cartAPI.Post("/", cart.PostCart)
	cartAPI.Get("/", cart.GetCart)
	cartAPI.Put("/:id", cart.PutCart)
	cartAPI.Get("/:id", cart.GetCartID)
	cartAPI.Delete("/:id", cart.DeleteCart)

	// Category
	categoryAPI := api.Group("/categories")
	categoryAPI.Use(middleware.Oauth2Authentication)
	categoryAPI.Post("/", category.PostCategory)
	categoryAPI.Get("/", category.GetCategory)
	categoryAPI.Put("/:id", category.PutCategory)
	categoryAPI.Get("/:id", category.GetCategoryID)
	categoryAPI.Delete("/:id", category.DeleteCategory)

	// Product
	productAPI := api.Group("/products")
	productAPI.Use(middleware.Oauth2Authentication)
	productAPI.Post("/", product.PostProduct)
	productAPI.Get("/", product.GetProduct)
	productAPI.Put("/:id", product.PutProduct)
	productAPI.Get("/:id", product.GetProductID)
	productAPI.Delete("/:id", product.DeleteProduct)

	// transaction
	transactionAPI := api.Group("/transactions")
	transactionAPI.Use(middleware.Oauth2Authentication)
	transactionAPI.Post("/", transaction.PostTransaction)
	transactionAPI.Post("/:id/payment", transaction.PostTransactionPayment)
	transactionAPI.Post("/:id/cancel", transaction.PostTransactionCancel)

	// User
	userAPI := api.Group("/users")
	userAPI.Use(middleware.Oauth2Authentication)
	userAPI.Get("/", user.GetUser)
	userAPI.Put("/:id", user.PutUser)
	userAPI.Get("/:id", user.GetUserID)

}
