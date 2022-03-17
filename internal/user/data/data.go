package data

import (
	"api-training-go/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db := openDB(c, logger)
	if db == nil {
		panic("db error")
	}
	return &Data{db}, cleanup, nil
}

func openDB(c *conf.Data, logger log.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Fatal("OPEN DB error")
		return nil
	}
	if c.Database.Debug {
		db = db.Debug()
	}
	return db
}
