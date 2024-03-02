package product

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

func TestPutProduct(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	app.Use(middleware.TokenValidator())

	app.Put("/products/:id", PutProduct)

	initial := model.Product{
		ProductAPI: model.ProductAPI{
			SKU:         nil,
			Name:        nil,
			Description: nil,
			Price:       nil,
			CategoryID:  nil,
		},
	}

	initial2 := model.Product{
		ProductAPI: model.ProductAPI{
			SKU:         nil,
			Name:        nil,
			Description: nil,
			Price:       nil,
			CategoryID:  nil,
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/products/" + initial.ID.String()

	payload := `{
		"sku": null,
		"name": null,
		"description": null,
		"price": null,
		"categoryid": null
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
	uri = "/products/non-existing-id"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test duplicate data
	uri = "/products/" + initial2.ID.String()
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 409, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
