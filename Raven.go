package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type user struct {
	Id       int64  `gorm:"column:id"`
	Username string `gorm:"size:150,column:username"`
}

func (user) TableName() string {
	return "auth_user"
}

func main() {
	db, err := gorm.Open("sqlite3", "db.sqlite")
	defer db.Close()
	user := &UserModel{}
	db.First(&user, 10)
	fmt.Println(user)
}
