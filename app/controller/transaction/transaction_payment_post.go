package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func PostTransactionPayment(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	id := c.Params("transaction_id")

	api := new(model.TransactionPaymentAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	var data model.TransactionPayment
	lib.Merge(api, &data)
	data.TransactionID = lib.StringToUUID(id)

	if err := db.Create(&data); err != nil {
		return lib.ErrorConflict(c, err)
	}

	amount := float64(0)
	db.Model(&model.TransactionPayment{}).Select(`sum(paid_amount) as amount`).Where(`transaction_id = ?`, id).Find(&amount)

	trx := model.Transaction{}
	db.Where(`id = ?`).Take(&trx)

	if amount >= lib.RevFloat64(trx.TotalPrice) {
		trx.TransactionStatus = lib.Strptr("paid")
	}
	return lib.OK(c, data)
}
