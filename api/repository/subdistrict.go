package repository

import (
	"api/entity"
	"api/helper"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

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
