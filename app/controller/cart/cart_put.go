package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutCart godoc
// @Summary Update Cart by id
// @Description Update Cart by id
// @Param id path string true "Cart ID"
// @Param data body model.CartAPI true "Cart data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Cart "Cart data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /carts/{id} [put]
// @Tags Cart
func PutCart(c *fiber.Ctx) error {
	api := new(model.CartPayload)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Cart
	result := db.Model(&data).
		Where(db.Where(model.Cart{
			Base: model.Base{
				ID: &id,
			},
		})).
		Preload("Product").
		Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	if lib.RevInt64(api.Quantity) != lib.RevInt64(data.Quantity) {
		product := data.Product
		data.Price = lib.Float64ptr(float64(lib.RevInt64(api.Quantity)) * lib.RevFloat64(product.Price))
		data.Quantity = api.Quantity
	}

	data.ModifierID = lib.GetXUserID(c)
	data.Notes = api.Notes

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err.Error())
	}

	return lib.OK(c, data)
}
