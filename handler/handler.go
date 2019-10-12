package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/letanthang/go_fw/types"
)

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetAllStudent(c echo.Context) error {
	aStudent := types.Student{ID: 1, FirstName: "Thai", LastName: "Le", Age: 80}
	bStudent := types.Student{ID: 2, FirstName: "Van", LastName: "Nguyen", Age: 90}
	cStudent := types.Student{ID: 3, FirstName: "Cuong", LastName: "Nguyen", Age: 50}
	dStudent := types.Student{ID: 4, FirstName: "Phat", LastName: "Le", Age: 100}

	result := []types.Student{aStudent, bStudent, cStudent, dStudent}

	// return c.String(http.StatusOK, "OK")
	return c.JSON(http.StatusOK, result)
}
