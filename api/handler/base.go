package handler

import (
	"api/entity"
	"api/usecase"
	"encoding/json"
	"net/http"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type MetaResponse struct {
	HTTPCode int `json:"http_code"`
	Limit    int `json:"limit,omitempty"`
	Offset   int `json:"offset,omitempty"`
	Total    int `json:"total,omitempty"`
}

type Response struct {
	Message string                     `json:"message,omitempty"`
	Data    interface{}                `json:"data,omitempty"`
	Errors  []entity.HttpResponseError `json:"errors,omitempty"`
	Meta    MetaResponse               `json:"meta"`
}

func (h *Handler) writeSuccess(w http.ResponseWriter, data interface{}, meta MetaResponse) {
	res := Response{
		Data: data,
		Meta: meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(meta.HTTPCode)
	responseBody, _ := json.Marshal(res)
	_, _ = w.Write(responseBody)
}

func (h *Handler) translateError(err error) (int, []entity.HttpResponseError) {
	switch origErr := err.(type) {
	case entity.HttpError:
		return origErr.ToHttpError()
	default:
		return entity.InternalServerError{Message: err.Error()}.ToHttpError()
	}
}

func (h *Handler) writeError(w http.ResponseWriter, err error) {
	statusCode, httpErrors := h.translateError(err)
	meta := MetaResponse{
		HTTPCode: statusCode,
	}
	res := Response{
		Errors: httpErrors,
		Meta: MetaResponse{
			HTTPCode: statusCode,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(meta.HTTPCode)
	responseBody, _ := json.Marshal(res)
	_, _ = w.Write(responseBody)
}
