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
	var isPay, isDetails bool = false, false

	api := new(model.TransactionPayload)
	if err := c.BodyParser(&api); nil != err {
		return lib.ErrorBadRequest(c)
	}

	trx := model.Transaction{}
	lib.Merge(api.Data, &trx)
	trx.ID = transactionID
	trx.CreatorID = userID
	trx.TransactionStatus = lib.Strptr("unpaid")

	payment := model.TransactionPayment{}
	if api.Payment != nil {
		lib.Merge(api.Payment, &payment)
		payment.TransactionID = transactionID
		if lib.RevFloat64(payment.PaidAmount) >= lib.RevFloat64(trx.TotalPrice) {
			trx.TransactionStatus = lib.Strptr("paid")
		}
		isPay = true
	}

	details := []model.TransactionDetail{}
	if api.Details != nil {
		for _, detail := range *api.Details {
			single := model.TransactionDetail{}
			lib.Merge(detail, &single)
			single.TransactionID = transactionID

			details = append(details, single)
		}
		isDetails = true
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		// do tx
		if err := tx.Create(&trx).Error; err != nil {
			return err
		}
		if isPay {
			if err := tx.Create(&payment).Error; err != nil {
				return err
			}
		}
		if isDetails {
			if err := tx.Create(&details).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if nil != err {
		return lib.ErrorInternal(c, err.Error())
	}

	return lib.OK(c)
}
