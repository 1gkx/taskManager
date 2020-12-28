package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var x *gorm.DB

func Initialize() error {
	var err error

	if x, err = gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{}); err != nil {
		return err
	}

	// if conf.Cfg.Database.Driver == "sqlite3" {
	// 	// x, err = gorm.Open(conf.Cfg.Database.Driver, conf.Cfg.Database.Path)
	// 	if x, err = gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{}); err != nil {
	// 		return err
	// 	}
	// }

	migrate()

	return nil
}

func GetEnginie() *gorm.DB {
	return x
}
