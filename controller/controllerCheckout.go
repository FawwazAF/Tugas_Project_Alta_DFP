package controller

import (
	"alta/project/database"
	"alta/project/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Checkout(c echo.Context) error {
	//Check Authentication
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loggedInUserId := middleware.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized access, you can only see your own")
	}

	// Checkout
	total, err := database.GetTotal()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":    "Checkout Success",
		"total_qty":   total.Total_qty,
		"total_price": total.Total_price,
	})
}
