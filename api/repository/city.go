package repository

import (
	"api/entity"
	"api/helper"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

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
