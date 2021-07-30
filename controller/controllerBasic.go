package controller

import (
	"alta/project/database"
	_ "alta/project/middleware"
	"alta/project/model"
	"net/http"

	"github.com/labstack/echo"
)

func DummyController(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)
	users, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     users,
	})
}

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

func NewItem(c echo.Context) error {
	// binding data
	Product := model.Product{}
	c.Bind(&Product)
	products, err := database.InputItem(Product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     products,
	})
}
