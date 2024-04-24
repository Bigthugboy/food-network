package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/bigthugboy/food-network/pkg/model"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InsertCustomer inserts a new customer into the database
func (f *foodNetwork) InsertCustomer(customer model.Customer) (int64, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if db == nil {
		return -1, fmt.Errorf("database connection is not initialized")
	}
	var existingCustomer model.Customer
	if err := db.Where("email = ?", customer.Email).First(&existingCustomer).Error; err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	if existingCustomer.ID != 0 {
		return -1, fmt.Errorf("user with email '%s' already exists", customer.Email)
	}
	result := db.Create(&customer)
	if err := result.Error; err != nil {
		return -1, err
	}
	// Return the number of rows affected
	return result.RowsAffected, nil
}
