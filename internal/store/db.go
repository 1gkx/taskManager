package store

import (
	// "fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/1gkx/taskmanager/internal/conf"
)

var x *gorm.DB

func Initialize() error {
	var err error

	if conf.Cfg.Database.Driver == "sqlite3" {
		x, err = gorm.Open(conf.Cfg.Database.Driver, conf.Cfg.Database.Path)
		if err != nil {
			return err
		}
	}

	migrate()

	return nil
}

func GetEnginie() *gorm.DB {
	return x
}
