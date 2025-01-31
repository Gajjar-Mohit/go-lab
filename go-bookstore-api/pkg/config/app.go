package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
    dsn := "root:mysecretpassword@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
    d, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDB() *gorm.DB {
    return db
}