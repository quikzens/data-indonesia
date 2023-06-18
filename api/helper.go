package main

import (
	"net/url"
	"strconv"
)

type OrderQueryParam struct {
	SortBy  string
	OrderBy string
	Default string
}

func BuildOrderQuery(param OrderQueryParam) string {
	orderQuery := "created_at DESC"

	if param.Default != "" {
		orderQuery = param.Default
	}

	if param.SortBy != "" {
		orderQuery = param.SortBy + " DESC"
		if param.OrderBy == "asc" {
			orderQuery = param.SortBy + " ASC"
		}
	}

	return orderQuery
}

func GetIntParam(queryParams url.Values, fieldName string, errMessage string) (int, error) {
	var value int
	if queryParams.Get(fieldName) != "" {
		var convertErr error
		value, convertErr = strconv.Atoi(queryParams.Get(fieldName))
		if convertErr != nil {
			return value, BadRequestError{Field: fieldName, Message: errMessage}
		}
	}
	return value, nil
}
