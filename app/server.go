package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", Test)
	e.GET("/users", AllUsers)
	e.POST("/users", NewUser)

	e.Logger.Panic(e.Start(":9300"))
}

func Test(c echo.Context) error {

	type Message struct{
		Message string `json:"message"`
	}

	msg := &Message{Message:"I'm Alive"}

	return c.JSON(http.StatusOK, msg)
}
