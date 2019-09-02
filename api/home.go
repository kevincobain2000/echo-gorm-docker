package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	fmt.Println("request")
	return c.String(http.StatusOK, "Main Server")
}
