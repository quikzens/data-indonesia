package main

import (
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
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

func GetQueryIntParam(r *http.Request, fieldName string, errMessage string) (int, error) {
	queryParams := r.URL.Query()
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

func GetUrlIntParam(r *http.Request, fieldName string, errMessage string) (int, error) {
	valueString := chi.URLParam(r, fieldName)
	value, err := strconv.ParseInt(valueString, 10, 32)
	if err != nil {
		return 0, BadRequestError{Field: fieldName, Message: errMessage}
	}
	return int(value), nil
}

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got %v; want %v", actual, expected)
	}
}

func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got %q; expected to contain: %q", actual, expectedSubstring)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}
