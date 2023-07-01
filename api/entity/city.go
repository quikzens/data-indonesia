package entity

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type City struct {
	Id         int            `json:"id"`
	ProvinceId int            `json:"province_id"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

func (City) TableName() string {
	return "cities"
}

type PaginateCitiesParam struct {
	BasePaginateParam
	ProvinceId int
}

type CityNotFoundError struct{}

func (e CityNotFoundError) Error() string {
	return "City Not Found"
}

func (e CityNotFoundError) ToHttpError() (int, []HttpResponseError) {
	return http.StatusNotFound, []HttpResponseError{
		{
			Field:   "City",
			Message: e.Error(),
		},
	}
}
