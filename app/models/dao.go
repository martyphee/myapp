package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	*gorm.DB
}

var DAO Dao

func InitDB() {
	DB, err := gorm.Open("postgres", "user=mphee dbname=gorm sslmode=disable")
	DB.LogMode(true)
	DB.LogMode(false)

	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
	}

	DB.DB()
	DB.CreateTable(&User{})

	DB.LogMode(true)
	DB.DB().Ping()
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.AutoMigrate(&User{})

	DAO = Dao{&DB}
}

func (db *Dao) GetFirst() *User {
	var user User

	if db.DB == nil {
		panic("db is nil")
	}

	// Gorm will know to use table "users" ("user" if pluralisation has been disabled) for all operations.
	db.DB.First(&user)

	return &user
}
