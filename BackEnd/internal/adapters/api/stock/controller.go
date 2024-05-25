package api

import "github.com/labstack/echo"

func init() {
	e := echo.New()
	e.Get("/sock", GetAllStock())
}
