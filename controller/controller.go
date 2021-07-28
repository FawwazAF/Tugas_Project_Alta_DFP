package controller

import (
	"alta/project/database"
	_ "alta/project/middleware"
	_ "alta/project/model"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetManyController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
