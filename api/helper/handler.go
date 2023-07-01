package helper

import (
	"api/entity"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetQueryIntParam(r *http.Request, fieldName string, errMessage string) (int, error) {
	queryParams := r.URL.Query()
	var value int
	if queryParams.Get(fieldName) != "" {
		var convertErr error
		value, convertErr = strconv.Atoi(queryParams.Get(fieldName))
		if convertErr != nil {
			return value, entity.BadRequestError{Field: fieldName, Message: errMessage}
		}
	}
	return value, nil
}

func GetUrlIntParam(r *http.Request, fieldName string, errMessage string) (int, error) {
	valueString := chi.URLParam(r, fieldName)
	value, err := strconv.ParseInt(valueString, 10, 32)
	if err != nil {
		return 0, entity.BadRequestError{Field: fieldName, Message: errMessage}
	}
	return int(value), nil
}
