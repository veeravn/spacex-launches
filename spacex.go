package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/launchesPast", getLaunches)
    e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(":8080"))
}
