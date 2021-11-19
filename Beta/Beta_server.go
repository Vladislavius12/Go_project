package main

import (
	"github.com/labstack/echo/v4"
	"net/http"

)

func main()  {
	// Echo instance
	eh := echo.New()
	eh.GET("/", func(a echo.Context) error{
		return a.String(http.StatusOK, "It's sh_t works!")
	})
	// Start server
	eh.Logger.Fatal(eh.Start(":1323"))
	eh.GET("/user/:id", getUser)
}
func getUser(a echo.Context) error{
	id := a.Param("id")
	return a.String(http.StatusOK, id)
}