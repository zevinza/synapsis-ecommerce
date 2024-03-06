package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostCart godoc
// @Summary Create new Cart
// @Description Create new Cart
// @Param data body model.CartPayload true "Cart data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Cart "Cart data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /carts [post]
// @Tags Cart
func PostCart(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	api := new(model.CartPayload)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	userID := lib.GetXUserID(c)
	if api.Quantity == nil {
		api.Quantity = lib.Int64ptr(1)
	}

	product := model.Product{}
	db.Where(`id = ?`, api.ProductID).Take(&product)

	var data model.Cart
	if row := db.Where(`product_id = ? AND user_id = ?`, api.ProductID, userID).Take(&data); row.RowsAffected > 0 {
		qty := lib.RevInt64(data.Quantity) + lib.RevInt64(api.Quantity)
		data.Quantity = &qty
		data.Price = lib.Float64ptr(float64(qty) * lib.RevFloat64(product.Price))
		if err := db.Updates(&data).Error; nil != err {
			return lib.ErrorConflict(c)
		}
		return lib.OK(c, data)
	}
	lib.Merge(api, &data)
	data.CreatorID = userID
	data.UserID = userID
	data.Price = lib.Float64ptr(float64(lib.RevInt64(api.Quantity)) * lib.RevFloat64(product.Price))

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err.Error())
	}
	data.Product = &product

	return lib.Created(c, data)
}
