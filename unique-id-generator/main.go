package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)
const (
	defaultLength = 36
)

type errorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
type successResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	UID     string `json:"uid"`
}

func main() {
	e := echo.New()
	e.GET("/", generateUIDHandler)
	e.Logger.Fatal(e.Start(":1324"))
}

type conditions struct {
	HasHyphen bool
	Length    int
}

func generateUIDHandler(c echo.Context) error {
	conditions := conditions{
		HasHyphen: true,
		Length:    defaultLength,
	}
	switch c.QueryParam("allowHyphen") {
	case "true", "":
		conditions.HasHyphen = true
	case "false":
		conditions.HasHyphen = false
	default:
		return c.JSON(http.StatusBadRequest, errorResponse{
			Status:  false,
			Message: "Invalid hyphen condition",
			Error:   "Hyphen can only be 'allowed' or 'disallowed'",
		})
	}
	length := c.QueryParam("len")
	if length != "" {
		lenInt64, err := strconv.ParseInt(length, 10, 36)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse{
				Status:  false,
				Message: "Invalid length",
				Error:   err.Error(),
			})
		}
		conditions.Length = int(lenInt64)
		switch {
		case conditions.Length < 0:
			return c.JSON(http.StatusBadRequest, errorResponse{
				Status:  false,
				Message: "Invalid length",
				Error:   "Length cannot be negative",
			})
		}
	}
	uid := uuid.NewString()

	for {
		if !conditions.HasHyphen {
			uid = strings.Replace(uid, "-", "", -1)
		}
		if len(uid) >= conditions.Length {
			break
		}
		uid += uuid.NewString()
	}
	uid = uid[:conditions.Length]
	if uid[conditions.Length-1] == '-' {
		uid = uid[:conditions.Length-1] + uuid.NewString()[0:1]
	}

	return c.JSON(http.StatusOK, successResponse{
		Status:  true,
		Message: "Unique ID generated successfully",
		UID:     uid,
	})
}
