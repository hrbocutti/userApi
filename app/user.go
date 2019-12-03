package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type User struct {
	Id int `gorm:"column:id; PRIMARY_KEY" json:"id"`
	FkVenture int `json:"venture"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Isactive bool `gorm:"column:is_active" json:"is_active"`
	CreateAt time.Time `gorm:"column:created_at"`
	UpdateAt time.Time `gorm:"column:updated_at"`
}

func AllUsers(c echo.Context) error {
	db := dbConnector()
	defer db.Close()
	var users []User
	db.SingularTable(true)
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func NewUser(c echo.Context ) (err error)  {
	u := new(User)

	if err = c.Bind(u); err != nil {
		return
	}
	db := dbConnector()
	db.SingularTable(true)
	db.Create(&u)
	db.Close()

	return c.JSON(http.StatusOK, 200)
}

func DeleteUser(w http.ResponseWriter, r *http.Request )  {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request )  {
	fmt.Fprintf(w, "Update Endpoint Hit")
}


func dbConnector() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error")
	}
	return db
}