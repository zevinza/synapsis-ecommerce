package migrations

import "api/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	model.Cart{},
	model.Category{},
	model.Product{},
	model.ReferenceCount{},
	model.Transaction{},
	model.TransactionDetail{},
	model.TransactionPayment{},
	model.User{},
}
