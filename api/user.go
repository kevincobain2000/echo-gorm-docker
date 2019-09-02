package api

import (
	"net/http"

	"github.com/kevincobain2000/echo-gorm-docker/db"
	"github.com/kevincobain2000/echo-gorm-docker/model"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}
	db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, users)
}
