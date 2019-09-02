package api

import (
	"math/rand"
	"net/http"

	"github.com/kevincobain2000/echo-gorm-docker/db"
	"github.com/kevincobain2000/echo-gorm-docker/model"
	"github.com/labstack/echo"
)

func CountItems(c echo.Context) error {
	db := db.DbManager()
	count := 0
	db.Model(model.Item{}).Count(&count)
	return c.JSON(http.StatusOK, count)
}
func Find(c echo.Context) error {
	db := db.DbManager()
	items := []model.Item{}
	randId := randomInt(1, 999999)
	db.Where("id = ?", randId).First(&items)
	return c.JSON(http.StatusOK, items)
}

func WhereIn(c echo.Context) error {
	db := db.DbManager()
	items := []model.Item{}
	ids := []int{randomInt(1000, 9999), randomInt(10000, 99999), randomInt(100000, 999999)}
	db.Where("id IN (?) ", ids).Find(&items)
	return c.JSON(http.StatusOK, items)
}

func WhereLike(c echo.Context) error {
	db := db.DbManager()
	items := []model.Item{}
	randomName := randomString(5)
	db.Where("name LIKE ? ", randomName+"%").Find(&items)
	return c.JSON(http.StatusOK, items)
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
