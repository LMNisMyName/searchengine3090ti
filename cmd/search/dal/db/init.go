package db

import (
	"searchengine3090ti/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
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
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
	// tables, _ := DB.Migrator().GetTables()
	// if len(tables) == 0 {
	// 	CreateTable()
	// }
	DropAllTable()
	CreateTable()
}

func DelectAllEntry() {
	DB.Where("1 = 1").Delete(&Keyword{})
	DB.Where("1 = 1").Delete(&UserID{})
	DB.Where("1 = 1").Delete(&Record{})
}

func DropAllTable() {
	tables, _ := DB.Migrator().GetTables()
	for _, table := range tables {
		DB.Migrator().DropTable(table)
	}
}

func CreateTable() {
	DB.AutoMigrate(&Keyword{})
	DB.AutoMigrate(&Record{})
}
