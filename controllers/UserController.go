package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test-api/assemblers"
	"test-api/services"
	"time"
)

type PostResponse struct {
	Status string `json:"status"`
	From   string `json:"from"`
	To     string `json:"to"`
}

func InitRoots(e *echo.Echo) {
	e.GET("/data", getUsers)
	e.POST("/data", echoPost)
}

func getUsers(c echo.Context) error {
	from := c.QueryParam("from")
	to := c.QueryParam("to")

	users := services.GetUsers()
	if from != "" && to != "" {
		timeFrom, err := time.Parse("2006-01-02", from)
		timeTo, err1 := time.Parse("2006-01-02", to)

		if err != nil || err1 != nil {
			return c.String(http.StatusBadRequest, "bad query param")
		}

		users = services.FilterUsers(timeFrom, timeTo, users)
	}

	response := assemblers.FormatUsersDataDto(users)
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func echoPost(c echo.Context) error {
	from := c.FormValue("from")
	to := c.FormValue("to")

	response := PostResponse{
		Status: "Success",
		From:   from,
		To:     to,
	}
	return c.JSONPretty(http.StatusOK, response, "  ")
}
