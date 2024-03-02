package cart

import (
	"api/app/config"
	"api/app/lib"
	"api/app/middleware"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/spf13/viper"
)

func TestPutCart(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	app.Use(middleware.TokenValidator())

	app.Put("/carts/:id", PutCart)

	initial := model.Cart{
		CartAPI: model.CartAPI{
			UserID:    nil,
			ProductID: nil,
			Quantity:  nil,
			Price:     nil,
		},
	}

	initial2 := model.Cart{
		CartAPI: model.CartAPI{
			UserID:    nil,
			ProductID: nil,
			Quantity:  nil,
			Price:     nil,
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/carts/" + initial.ID.String()

	payload := `{
		"userid": null,
		"productid": null,
		"quantity": null,
		"price": null
	}`

	headers := map[string]string{
		"Content-Type":                      "application/json",
		viper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
	}

	response, body, err := lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")

	// test invalid json body
	response, _, err = lib.PutTest(app, uri, headers, "invalid json format")
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, response.StatusCode, "getting response code")

	// test update with non existing id
	uri = "/carts/non-existing-id"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test duplicate data
	uri = "/carts/" + initial2.ID.String()
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 409, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
