package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// OK is struct healthcheck
type OK struct {
	Status  int
	Message string
}

func main() {
	e := echo.New()
	e.GET("/", healthcheck)
	listenPort := ":5000"
	e.Logger.Fatal(e.Start(listenPort))
	fmt.Println("listen on : http://localhost" + listenPort)
}

func healthcheck(c echo.Context) error {
	m := OK{200, "Hello"}
	return c.JSON(http.StatusOK, m)
}
