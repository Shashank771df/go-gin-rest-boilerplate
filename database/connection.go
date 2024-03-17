package database

import (
	"fmt"

	"go-gin-rest-boilerplate/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"

var db *gorm.DB
var err error

/*
dbType can be 'MySql', 'Postrges', ”
*/
func Init() {
	////MySQL
	dsn := "root:@/go-rest-boilerplate?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//PostgreSQL
	//db, err = gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")

	////SQLite3
	//db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	//
	////SQL Server
	//db, err = gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")

	if err != nil {
		fmt.Println(err)
	}

	//db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.User{})
}

func GetDb() *gorm.DB {
	return db
}

// func CloseDb() {
// 	gorm.DB.
// }
