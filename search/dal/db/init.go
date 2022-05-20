package db

import (
	"search/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

//建立数据模型
//Keyword与UserID之间是manytomany关系，UserID与Record是has one关系
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
	// m := DB.Migrator()

	// if !m.HasTable(&Keyword{}) {
	// 	if err = m.CreateTable(&Keyword{}); err != nil {
	// 		panic(err)
	// 	}
	// }

	// if !m.HasTable(&UserID{}) {
	// 	if err = m.CreateTable(&UserID{}); err != nil {
	// 		panic(err)
	// 	}
	// }

	// if !m.HasTable(&Record{}) {
	// 	if err = m.CreateTable(&Record{}); err != nil {
	// 		panic(err)
	// 	}
	// }
	DB.AutoMigrate(&Keyword{})
	DB.AutoMigrate(&Record{})
}
