package entity

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Village struct {
	Id            int            `json:"id"`
	ProvinceId    int            `json:"province_id"`
	CityId        int            `json:"city_id"`
	SubdistrictId int            `json:"subdistrict_id"`
	Name          string         `json:"name"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

func (Village) TableName() string {
	return "villages"
}

type PaginateVillagesParam struct {
	BasePaginateParam
	ProvinceId    int
	CityId        int
	SubdistrictId int
}

type VillageNotFoundError struct{}

func (e VillageNotFoundError) Error() string {
	return "Village Not Found"
}

func (e VillageNotFoundError) ToHttpError() (int, []HttpResponseError) {
	return http.StatusNotFound, []HttpResponseError{
		{
			Field:   "Village",
			Message: e.Error(),
		},
	}
}
