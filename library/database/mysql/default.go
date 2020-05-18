package mysql

import (
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB(config *Config) *gorm.DB {
	db, err := gorm.Open("mysql", config.Dsn)
	if err != nil {
		panic(err)
	}

	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	return db
}
