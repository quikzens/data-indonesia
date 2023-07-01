package usecase

import (
	"api/entity"
	"context"
)

func (u *Usecase) GetCities(ctx context.Context, param entity.PaginateCitiesParam) (entity.PaginateResult[entity.City], error) {
	cities, count, err := u.repo.SearchCities(ctx, param)
	if err != nil {
		return entity.PaginateResult[entity.City]{}, err
	}

	return entity.PaginateResult[entity.City]{
		Data:   cities,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) PaginateCities(ctx context.Context, param entity.PaginateCitiesParam) (entity.PaginateResult[entity.City], error) {
	if param.Limit == 0 {
		param.Limit = 10 // fallback to default 10 paginate
	}

	cities, count, err := u.repo.SearchCities(ctx, param)
	if err != nil {
		return entity.PaginateResult[entity.City]{}, err
	}

	return entity.PaginateResult[entity.City]{
		Data:   cities,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) GetCity(ctx context.Context, id int) (entity.City, error) {
	return u.repo.FindCityByID(ctx, id)
}
