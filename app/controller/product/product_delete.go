package product

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteProduct godoc
// @Summary Delete Product by id
// @Description Delete Product by id
// @Param id path string true "Product ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /products/{id} [delete]
// @Tags Product
func DeleteProduct(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var data model.Product
	result := db.Model(&data).Where("id = ?", c.Params("id")).Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Delete(&data)

	return lib.OK(c)
}
