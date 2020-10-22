package main

import (
	"github.com/labstack/echo/v4"
	"test-api/controllers"
)

func main() {
	e := echo.New()

	controllers.InitRoots(e)

	e.Logger.Fatal(e.Start(":8080"))
}
