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
// @Param data body model.CartAPI true "Cart data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Cart "Cart data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security TokenKey
// @Router /carts [post]
// @Tags Cart
func PostCart(c *fiber.Ctx) error {
	api := new(model.CartAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())

	var data model.Cart
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err.Error())
	}

	return lib.Created(c, data)
}
