package product

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetProductCategory godoc
// @Summary List of Product by category ID
// @Description List of Product by category ID
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Param id path string true "Category ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.Product} "List of Product"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /products/category/{id} [get]
// @Tags Product
func GetProductCategory(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext()).WithContext(c.UserContext())
	pg := services.PG
	id, _ := uuid.Parse(c.Params("id"))

	mod := db.Model(&model.Product{}).Where(`category_id = ?`, id).Preload("Category")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Product{})

	return lib.OK(c, page)
}
