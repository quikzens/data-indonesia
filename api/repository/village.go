package repository

import (
	"api/entity"
	"api/helper"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

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
