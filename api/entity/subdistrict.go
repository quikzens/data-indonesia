package entity

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Subdistrict struct {
	Id         int            `json:"id"`
	ProvinceId int            `json:"province_id"`
	CityId     int            `json:"city_id"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

func (Subdistrict) TableName() string {
	return "subdistricts"
}

type PaginateSubdistrictsParam struct {
	BasePaginateParam
	ProvinceId int
	CityId     int
}

type SubdistrictNotFoundError struct{}

func (e SubdistrictNotFoundError) Error() string {
	return "Subdistrict Not Found"
}

func (e SubdistrictNotFoundError) ToHttpError() (int, []HttpResponseError) {
	return http.StatusNotFound, []HttpResponseError{
		{
			Field:   "Subdistrict",
			Message: e.Error(),
		},
	}
}
