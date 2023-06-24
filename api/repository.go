package main

import (
	"context"
	"errors"
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

func (r *Repository) FindProvinceByID(ctx context.Context, id int) (Province, error) {
	var province Province
	err := r.db.WithContext(ctx).First(&province, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return province, ProvinceNotFoundError{}
		} else {
			return province, err
		}
	}

	return province, nil
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

func (r *Repository) FindCityByID(ctx context.Context, id int) (City, error) {
	var city City
	err := r.db.WithContext(ctx).First(&city, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return city, CityNotFoundError{}
		} else {
			return city, err
		}
	}

	return city, nil
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

func (r *Repository) FindSubdistrictByID(ctx context.Context, id int) (Subdistrict, error) {
	var subdistrict Subdistrict
	err := r.db.WithContext(ctx).First(&subdistrict, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return subdistrict, SubdistrictNotFoundError{}
		} else {
			return subdistrict, err
		}
	}

	return subdistrict, nil
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

func (r *Repository) FindVillageByID(ctx context.Context, id int) (Village, error) {
	var village Village
	err := r.db.WithContext(ctx).First(&village, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return village, VillageNotFoundError{}
		} else {
			return village, err
		}
	}

	return village, nil
}
