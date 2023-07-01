package repository

import (
	"api/entity"
	"api/helper"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

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
