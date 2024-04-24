package repo

import (
	"github.com/bigthugboy/food-network/db/query"
	"github.com/bigthugboy/food-network/pkg/config"

	"github.com/jinzhu/gorm"
)

type foodNetwork struct {
	App *config.AppTools
	DB  *gorm.DB
}

func NewFoodNetworkDB(app *config.AppTools, db *gorm.DB) query.DBstore {
	return &foodNetwork{
		App: app,
		DB:  db,
	}
}
