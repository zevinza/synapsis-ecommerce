package migrations

import "api/app/model"

var (
	category model.Category
	product  model.Product
	refCount model.ReferenceCount
	user     model.User
)

// DataSeeds data to seeds
func DataSeeds() []interface{} {
	return []interface{}{
		category.Seed(),
		product.Seed(),
		refCount.Seed(),
		user.Seed(),
	}
}
