package db

import (
	qr "github.com/qingyggg/aufer/biz/model/query"
	"github.com/qingyggg/aufer/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Init init DB
func Init() (error, *qr.Query) {
	var err error
	db, err := gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	query := qr.Use(db)
	if err != nil {
		return err, nil
	}
	return nil, query
}
