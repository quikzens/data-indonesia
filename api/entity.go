package main

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

type PaginateResult[T any] struct {
	Data   []T
	Total  int
	Limit  int
	Offset int
}

type BasePaginateParam struct {
	Keywords string
	SortBy   string
	OrderBy  string
	Limit    int
	Offset   int
}

type PaginateProvincesParam struct {
	BasePaginateParam
}

type PaginateCitiesParam struct {
	BasePaginateParam
	ProvinceId int
}

type PaginateSubdistrictsParam struct {
	BasePaginateParam
	ProvinceId int
	CityId     int
}

type PaginateVillagesParam struct {
	BasePaginateParam
	ProvinceId    int
	CityId        int
	SubdistrictId int
}

type GetTotalResult struct {
	Provinces    int `json:"provinces"`
	Cities       int `json:"cities"`
	Subdistricts int `json:"subdistricts"`
	Villages     int `json:"villages"`
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
