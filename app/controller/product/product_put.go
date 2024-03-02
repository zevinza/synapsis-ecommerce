package product

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutProduct godoc
// @Summary Update Product by id
// @Description Update Product by id
// @Param id path string true "Product ID"
// @Param data body model.ProductAPI true "Product data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Product "Product data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security TokenKey
// @Router /products/{id} [put]
// @Tags Product
func PutProduct(c *fiber.Ctx) error {
	api := new(model.ProductAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Product
	result := db.Model(&data).
		Where(db.Where(model.Product{
			Base: model.Base{
				ID: &id,
			},
		})).
		Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	lib.Merge(api, &data)
	data.ModifierID = lib.GetXUserID(c)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err.Error())
	}

	return lib.OK(c, data)
}
