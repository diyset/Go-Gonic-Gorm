package config

import (
	"fmt"
	"my-rest/structs"
	"os"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBinit() *gorm.DB {
	envResource := godotenv.Load()
	if envResource != nil {
		fmt.Println(envResource)
	}
	db_user := os.Getenv("db_user")
	db_pass := os.Getenv("db_pass")
	db_name := os.Getenv("db_name")
	db_host := os.Getenv("db_host")
	db_port := os.Getenv("db_port")

	// db_user:db_pass@tcp(db_host:db_port)/db_name?optional
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name) //Build connection string
	fmt.Println(dbURI)
	// db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar-golang?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		panic("Failed to connect to Database")
	}
	db.LogMode(true)
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
	// 	&structs.Nasabah{}, &structs.Alamat{},
	// )
	db.AutoMigrate(
		&structs.Nasabah{},
		&structs.Alamat{},
		&structs.Product{},
		&structs.Warehouse{},
		&structs.Order{},
		&structs.OrderDetail{},
		)
	return db
}
