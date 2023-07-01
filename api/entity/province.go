package entity

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Province struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (Province) TableName() string {
	return "provinces"
}

type PaginateProvincesParam struct {
	BasePaginateParam
}

type ProvinceNotFoundError struct{}

func (e ProvinceNotFoundError) Error() string {
	return "Province Not Found"
}

func (e ProvinceNotFoundError) ToHttpError() (int, []HttpResponseError) {
	return http.StatusNotFound, []HttpResponseError{
		{
			Field:   "Province",
			Message: e.Error(),
		},
	}
}
