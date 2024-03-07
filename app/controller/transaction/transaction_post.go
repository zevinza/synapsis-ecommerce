package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PostTransactionPayment godoc
// @Summary Post new Transaction
// @Description Post new Transaction
// @Param data body model.TransactionPayload true "Transaction data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Transaction "Transaction data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /transactions/ [post]
// @Tags Transaction
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
	lib.Merge(api, &trx)
	trx.ID = transactionID
	trx.CreatorID = userID
	trx.UserID = userID
	trx.TransactionStatus = lib.Strptr("unpaid")
	if trx.InvoiceNo == nil {
		trx.InvoiceNo = model.GenRefCount("Transaction", db)
	}

	user := model.User{}
	db.Where(`id = ?`, userID).Take(&user)
	trx.ContactName = lib.Strptr(lib.RevStr(user.FirstName) + " " + lib.RevStr(user.LastName))
	trx.ContactDetail = user.PhoneNumber
	if trx.ContactDetail == nil {
		trx.ContactDetail = user.Email
	}

	var total float64 = 0
	details := []model.TransactionDetail{}
	ids := []uuid.UUID{}
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
			if cart.ID != nil {
				ids = append(ids, *cart.ID)
			}
		}
		isDetails = true
	}
	trx.TotalProductPrice = lib.Float64ptr(total)
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

		if err := tx.Where(`id IN ?`, ids).Delete(&model.Cart{}).Error; err != nil {
			return err
		}

		return nil
	})
	if nil != err {
		return lib.ErrorInternal(c, err.Error())
	}

	return lib.OK(c, trx)
}
