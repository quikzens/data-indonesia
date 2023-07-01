package main

import (
	"api/entity"
	"api/helper"
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

func (r *Repository) SearchProvinces(ctx context.Context, param entity.PaginateProvincesParam) ([]entity.Province, int, error) {
	var count int64
	var res []entity.Province

	query := r.db.WithContext(ctx)
	if param.Keywords != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", param.Keywords))
	}

	err := query.Model(entity.Province{}).Count(&count).Error
	if err != nil {
		return res, int(count), err
	}

	orderQuery := helper.BuildOrderQuery(entity.OrderQueryParam{
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

func (r *Repository) FindProvinceByID(ctx context.Context, id int) (entity.Province, error) {
	var province entity.Province
	err := r.db.WithContext(ctx).First(&province, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return province, entity.ProvinceNotFoundError{}
		} else {
			return province, err
		}
	}

	return province, nil
}

func (r *Repository) SearchCities(ctx context.Context, param entity.PaginateCitiesParam) ([]entity.City, int, error) {
	var count int64
	var res []entity.City

	query := r.db.WithContext(ctx)
	if param.Keywords != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", param.Keywords))
	}

	if param.ProvinceId != 0 {
		query = query.Where("province_id = ?", param.ProvinceId)
	}

	err := query.Model(entity.City{}).Count(&count).Error
	if err != nil {
		return res, int(count), err
	}

	orderQuery := helper.BuildOrderQuery(entity.OrderQueryParam{
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

func (r *Repository) FindCityByID(ctx context.Context, id int) (entity.City, error) {
	var city entity.City
	err := r.db.WithContext(ctx).First(&city, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return city, entity.CityNotFoundError{}
		} else {
			return city, err
		}
	}

	return city, nil
}

func (r *Repository) SearchSubdistricts(ctx context.Context, param entity.PaginateSubdistrictsParam) ([]entity.Subdistrict, int, error) {
	var count int64
	var res []entity.Subdistrict

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

	err := query.Model(entity.Subdistrict{}).Count(&count).Error
	if err != nil {
		return res, int(count), err
	}

	orderQuery := helper.BuildOrderQuery(entity.OrderQueryParam{
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

func (r *Repository) FindSubdistrictByID(ctx context.Context, id int) (entity.Subdistrict, error) {
	var subdistrict entity.Subdistrict
	err := r.db.WithContext(ctx).First(&subdistrict, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return subdistrict, entity.SubdistrictNotFoundError{}
		} else {
			return subdistrict, err
		}
	}

	return subdistrict, nil
}

func (r *Repository) SearchVillages(ctx context.Context, param entity.PaginateVillagesParam) ([]entity.Village, int, error) {
	var count int64
	var res []entity.Village

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

	err := query.Model(entity.Village{}).Count(&count).Error
	if err != nil {
		return res, int(count), err
	}

	orderQuery := helper.BuildOrderQuery(entity.OrderQueryParam{
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

func (r *Repository) FindVillageByID(ctx context.Context, id int) (entity.Village, error) {
	var village entity.Village
	err := r.db.WithContext(ctx).First(&village, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return village, entity.VillageNotFoundError{}
		} else {
			return village, err
		}
	}

	return village, nil
}
