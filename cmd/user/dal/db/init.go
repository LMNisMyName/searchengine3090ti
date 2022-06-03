package db

import (
	"searchengine3090ti/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return
	}
	if err = m.CreateTable(&User{}); err != nil {
		panic(err)
	}
}
