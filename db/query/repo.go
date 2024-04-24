package query

import (
	"github.com/bigthugboy/food-network/pkg/model"
)

type DBstore interface {
	InsertCustomer(model.Customer) (int64, error)
}
