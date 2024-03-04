package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PostTransaction(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	userID := lib.GetXUserID(c)
	transactionID := lib.GenUUID()
	var isDetails bool = false

	api := new(model.TransactionPayload)
	if err := c.BodyParser(&api); nil != err {
		return lib.ErrorBadRequest(c)
	}

	trx := model.Transaction{}
	trx.ID = transactionID
	trx.CreatorID = userID
	trx.TransactionStatus = lib.Strptr("unpaid")
	trx.Notes = api.Notes
	trx.Description = api.Description
	if trx.InvoiceNo == nil {
		trx.InvoiceNo = model.GenRefCount("Transaction", db)
	}

	var total float64 = 0
	details := []model.TransactionDetail{}
	if api.CartIds != nil {
		carts := []model.Cart{}
		db.Where(`id IN ?`, *api.CartIds).Preload("Product").Find(&carts)
		for _, cart := range carts {
			product := model.Product{}
			if cart.Product != nil {
				product = *cart.Product
			}
			price := lib.RevFloat64(product.Price) * float64(lib.RevInt64(cart.Quantity))
			detail := model.TransactionDetail{
				TransactionDetailAPI: model.TransactionDetailAPI{
					TransactionID:      transactionID,
					ProductID:          cart.ProductID,
					ProductName:        product.Name,
					ProductSKU:         product.SKU,
					ProductDescription: product.Description,
					ProductPrice:       product.Price,
					Quantity:           cart.Quantity,
					SubtotalPrice:      &price,
					Notes:              cart.Notes,
				},
			}
			details = append(details, detail)
			total += price
		}
		isDetails = true
	}
	total = total - lib.RevFloat64(api.Discount) + lib.RevFloat64(api.Fee)
	trx.TotalPrice = lib.Float64ptr(total)
	trx.TotalDiscount = api.Discount
	trx.TotalFee = api.Fee

	err := db.Transaction(func(tx *gorm.DB) error {
		// do tx
		if err := tx.Create(&trx).Error; err != nil {
			return err
		}

		if isDetails {
			if err := tx.Create(&details).Error; err != nil {
				return err
			}
			trx.Details = &details
		}

		return nil
	})
	if nil != err {
		return lib.ErrorInternal(c, err.Error())
	}

	return lib.OK(c, trx)
}
