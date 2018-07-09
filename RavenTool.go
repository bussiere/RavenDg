package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// UserModel is a representation of a django User
type UserModel struct {
	ID       int64  `gorm:"column:id"`
	Username string `gorm:"size:150,column:username"`
	Password string `gorm:"size:150,column:username"`
}

// TableName (UserModel) is a function to rename user to auth_user
func (UserModel) TableName() string {
	return "auth_user"
}

// BoxCase is a representation of a django User
type BoxCase struct {
	ID       int64  `gorm:"column:id"`
	Username string `gorm:"size:150,column:username"`
	Password string `gorm:"size:150,column:username"`
}

// TableName (BoxCase) is a function to rename user to auth_user
func (BoxCase) TableName() string {
	return "Box_case"
}

func ping(c *gin.Context) {
	user := &UserModel{}
	db.First(&user, 1)
	fmt.Println(user.Username)
	c.JSON(200, gin.H{
		"message": user.Username,
	})
}

var db *gorm.DB

func main() {
	db, _ = gorm.Open("sqlite3", "db.sqlite3")
	defer db.Close()
	user := &UserModel{}
	db.First(&user, 1)
	fmt.Println(user.Username)
	users := []UserModel{}
	db.Find(&users)
	for index, element := range users {
		fmt.Println(index)
		fmt.Println(element)
	}
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}
