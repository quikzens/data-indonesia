package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SearchProvinces(ctx context.Context, param PaginateProvincesParam) ([]Province, int, error) {
	var count int64
	var res []Province

	query := r.db.WithContext(ctx)
	if param.Keywords != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", param.Keywords))
	}

	err := query.Model(Province{}).Count(&count).Error
	if err != nil {
		return res, int(count), err
	}

	orderQuery := BuildOrderQuery(OrderQueryParam{
		SortBy:  param.SortBy,
		OrderBy: param.OrderBy,
	})

	if param.Limit > 0 {
		query = query.Limit(param.Limit)
	}

	err = query.Offset(param.Offset).Order(orderQuery).Find(&res).Error
	if err != nil {
		return res, int(count), err
	}

	return res, int(count), nil
}

func (r *Repository) SearchCities(ctx context.Context, param PaginateCitiesParam) ([]City, int, error) {
	var count int64
	var res []City

	query := r.db.WithContext(ctx)
	if param.Keywords != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", param.Keywords))
	}

	if param.ProvinceId != 0 {
		query = query.Where("province_id = ?", param.ProvinceId)
	}

	err := query.Model(City{}).Count(&count).Error
	if err != nil {
		return res, int(count), err
	}

	orderQuery := BuildOrderQuery(OrderQueryParam{
		SortBy:  param.SortBy,
		OrderBy: param.OrderBy,
	})

	if param.Limit > 0 {
		query = query.Limit(param.Limit)
	}

	err = query.Offset(param.Offset).Order(orderQuery).Find(&res).Error
	if err != nil {
		return res, int(count), err
	}

	return res, int(count), nil
}

func (r *Repository) SearchSubdistricts(ctx context.Context, param PaginateSubdistrictsParam) ([]Subdistrict, int, error) {
	var count int64
	var res []Subdistrict

	query := r.db.WithContext(ctx)
	if param.Keywords != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", param.Keywords))
	}

	if param.ProvinceId != 0 {
		query = query.Where("province_id = ?", param.ProvinceId)
	}

	if param.CityId != 0 {
		query = query.Where("city_id = ?", param.CityId)
	}

	err := query.Model(Subdistrict{}).Count(&count).Error
	if err != nil {
		return res, int(count), err
	}

	orderQuery := BuildOrderQuery(OrderQueryParam{
		SortBy:  param.SortBy,
		OrderBy: param.OrderBy,
	})

	if param.Limit > 0 {
		query = query.Limit(param.Limit)
	}

	err = query.Offset(param.Offset).Order(orderQuery).Find(&res).Error
	if err != nil {
		return res, int(count), err
	}

	return res, int(count), nil
}

func (r *Repository) SearchVillages(ctx context.Context, param PaginateVillagesParam) ([]Village, int, error) {
	var count int64
	var res []Village

	query := r.db.WithContext(ctx)
	if param.Keywords != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", param.Keywords))
	}

	if param.ProvinceId != 0 {
		query = query.Where("province_id = ?", param.ProvinceId)
	}

	if param.CityId != 0 {
		query = query.Where("city_id = ?", param.CityId)
	}

	if param.SubdistrictId != 0 {
		query = query.Where("subdistrict_id = ?", param.SubdistrictId)
	}

	err := query.Model(Village{}).Count(&count).Error
	if err != nil {
		return res, int(count), err
	}

	orderQuery := BuildOrderQuery(OrderQueryParam{
		SortBy:  param.SortBy,
		OrderBy: param.OrderBy,
	})

	if param.Limit > 0 {
		query = query.Limit(param.Limit)
	}

	err = query.Offset(param.Offset).Order(orderQuery).Find(&res).Error
	if err != nil {
		return res, int(count), err
	}

	return res, int(count), nil
}
