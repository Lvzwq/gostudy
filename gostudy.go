package main

import (
	"fmt"
	"gostudy/api"
)
import "github.com/labstack/echo"

func main() {
	fmt.Println("Hello World")
	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
