package usecase

import (
	"api/entity"
	"context"
)

func (u *Usecase) GetProvinces(ctx context.Context, param entity.PaginateProvincesParam) (entity.PaginateResult[entity.Province], error) {
	provinces, count, err := u.repo.SearchProvinces(ctx, param)
	if err != nil {
		return entity.PaginateResult[entity.Province]{}, err
	}

	return entity.PaginateResult[entity.Province]{
		Data:   provinces,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) PaginateProvinces(ctx context.Context, param entity.PaginateProvincesParam) (entity.PaginateResult[entity.Province], error) {
	if param.Limit == 0 {
		param.Limit = 10 // fallback to default 10 paginate
	}

	provinces, count, err := u.repo.SearchProvinces(ctx, param)
	if err != nil {
		return entity.PaginateResult[entity.Province]{}, err
	}

	return entity.PaginateResult[entity.Province]{
		Data:   provinces,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) GetProvince(ctx context.Context, id int) (entity.Province, error) {
	return u.repo.FindProvinceByID(ctx, id)
}
