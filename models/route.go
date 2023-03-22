package models

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	//"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB
func ConnectDatabase(){
	fmt.Println("Go MySQL Tutorial")
	
	Database, err := gorm.Open(mysql.Open("admin:admin@tcp(localhost:3306)/gin"))//err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gin")
	if err != nil{
		panic(err)
	}
	//Database.AutoMigrate()
	//Database.AutoMigrate()
	//defer Database.Close()
	Database.AutoMigrate(&Product{})
	DB = Database
}