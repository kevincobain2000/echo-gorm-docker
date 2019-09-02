package route

import (
	"github.com/kevincobain2000/echo-gorm-docker/api"

	"github.com/labstack/echo"
)

// Init it
func Init() *echo.Echo {
	e := echo.New()

	// e.GET("/", api.Home)
	// e.GET("/api/health/count", api.CountItems)
	e.GET("/", api.Home)
	e.GET("/api/health/find", api.Find)
	e.GET("/api/health/where-like", api.WhereLike)
	e.GET("/api/health/where-in", api.WhereIn)
	// e.GET("/users", api.GetUsers)
	return e
}
