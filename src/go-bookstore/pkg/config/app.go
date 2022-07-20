package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	// using gorm to open up a connection to the DB
	d, err := gorm.Open("mysql", "tjapit:testing321@/simplerest?charset=utf8&parseTime=True&loc=Local")
	// error check
	if err != nil {
		panic(err)
	}
	// assign database
	db = d
}

func GetDB() *gorm.DB {
	return db
}
