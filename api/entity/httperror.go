package entity

import (
	"net/http"
)

type HttpError interface {
	ToHttpError() (int, []HttpResponseError)
}

type HttpResponseError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type InternalServerError struct {
	Message string
}

func (e InternalServerError) Error() string {
	return e.Message
}

func (e InternalServerError) ToHttpError() (int, []HttpResponseError) {
	return http.StatusInternalServerError, []HttpResponseError{
		{
			Message: e.Message,
		},
	}
}

type BadRequestError struct {
	Field   string
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

func (e BadRequestError) ToHttpError() (int, []HttpResponseError) {
	return http.StatusBadRequest, []HttpResponseError{
		{
			Field:   e.Field,
			Message: e.Message,
		},
	}
}
