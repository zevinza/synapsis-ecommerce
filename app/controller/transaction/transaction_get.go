package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetTransaction godoc
// @Summary List of Transaction
// @Description List of Transaction
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.Transaction} "List of Transaction"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /transactions [get]
// @Tags Transaction
func GetTransaction(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext()).WithContext(c.UserContext())
	pg := services.PG

	mod := db.Model(&model.Transaction{}).Where(`user_id = ?`, lib.GetXUserID(c)).Preload("Details").Preload("Payments")


	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, page)
}
