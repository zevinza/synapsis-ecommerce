package category

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

func TestDeleteCategory(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	app.Use(middleware.TokenValidator())

	app.Delete("/categories/:id", DeleteCategory)

	initial := model.Category{
		CategoryAPI: model.CategoryAPI{
			Code: nil,
			Name: nil,
		},
	}

	db.Create(&initial)

	headers := map[string]string{
		viper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
	}

	uri := "/categories/" + initial.ID.String()
	response, _, err := lib.DeleteTest(app, uri, headers)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")

	// test delete with non existing id
	response, _, err = lib.DeleteTest(app, uri, headers)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test invalid token
	response, _, err = lib.DeleteTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 401, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
