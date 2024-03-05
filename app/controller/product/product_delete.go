package product

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	if !lib.GetXIsAdmin(c) {
		return lib.ErrorUnauthorized(c)
	}
	db := services.DB.WithContext(c.UserContext())
	id := c.Params("id")

	var data model.Product
	result := db.Model(&data).Where("id = ?", id).Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&data).Error; err != nil {
			return err
		}

		if err := tx.Where(`product_id = ?`, id).Delete(&model.Cart{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c)
}
