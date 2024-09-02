package main

import (
	"net/http"
	"strconv"

	"github.com/AbdulRahimOM/misc-projects/password-generator/password"
	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
type successResponse struct {
	Status   bool   `json:"status"`
	Message  string `json:"message"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()
	e.GET("/", generatePasswordHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func generatePasswordHandler(c echo.Context) error {
	conditions := password.Conditions{
		HasSmallAlph: true,
		HasCapAlph:   true,
		HasNum:       true,
		HasSymbol:    false,
		MinLength:    8,
		MaxLength:    18,
	}
	if c.QueryParam("smallAlph") == "false" {
		conditions.HasSmallAlph = false
		conditions.NumConditions++
	}
	if c.QueryParam("capAlph") == "true" {
		conditions.HasCapAlph = true
		conditions.NumConditions++
	}
	if c.QueryParam("num") == "true" {
		conditions.HasNum = true
		conditions.NumConditions++
	}
	if c.QueryParam("HasSymbol") == "true" {
		conditions.HasSymbol = true
		conditions.NumConditions++
	}
	minLen := c.QueryParam("minLen")
	if minLen != "" {
		minLen, err := strconv.ParseInt(minLen, 10, 36)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse{
				Status:  false,
				Message: "Invalid minLength",
				Error:   err.Error(),
			})
		}
		conditions.MinLength = int8(minLen)
		if conditions.MinLength < conditions.NumConditions {
			conditions.MinLength = conditions.NumConditions
		}
	}

	maxLen := c.QueryParam("maxLen")
	if maxLen == "" {
		conditions.MaxLength = conditions.MinLength + 10
	} else {
		maxLen, err := strconv.ParseUint(maxLen, 10, 36)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorResponse{
				Status:  false,
				Message: "Invalid maxLength",
				Error:   err.Error(),
			})
		}
		conditions.MaxLength = int8(maxLen)
		if conditions.MaxLength < conditions.MinLength {
			return c.JSON(http.StatusBadRequest, errorResponse{
				Status:  false,
				Message: "Max limit is lower than minLimit",
				Error:   "Invalid maxLength",
			})
		}
		if conditions.MaxLength < conditions.NumConditions {
			return c.JSON(http.StatusBadRequest, errorResponse{
				Status:  false,
				Message: "Num of conditions are greater than maxLength",
				Error:   "Invalid maxLength",
			})
		}
	}

	if conditions.NumConditions == 0 {
		//default
		conditions.HasSmallAlph = true
		conditions.HasCapAlph = true
		conditions.HasNum = true
		conditions.HasSymbol = true
		conditions.NumConditions = 4
	}

	password := password.CreatePassword(conditions)

	return c.JSON(http.StatusOK, successResponse{
		Status:   true,
		Message:  "Password generated successfully",
		Password: password,
	})
}
