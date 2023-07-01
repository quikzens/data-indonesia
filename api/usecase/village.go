package usecase

import (
	"api/entity"
	"context"
)

func (u *Usecase) GetVillages(ctx context.Context, param entity.PaginateVillagesParam) (entity.PaginateResult[entity.Village], error) {
	villages, count, err := u.repo.SearchVillages(ctx, param)
	if err != nil {
		return entity.PaginateResult[entity.Village]{}, err
	}

	return entity.PaginateResult[entity.Village]{
		Data:   villages,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) PaginateVillages(ctx context.Context, param entity.PaginateVillagesParam) (entity.PaginateResult[entity.Village], error) {
	if param.Limit == 0 {
		param.Limit = 10 // fallback to default 10 paginate
	}

	villages, count, err := u.repo.SearchVillages(ctx, param)
	if err != nil {
		return entity.PaginateResult[entity.Village]{}, err
	}

	return entity.PaginateResult[entity.Village]{
		Data:   villages,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) GetVillage(ctx context.Context, id int) (entity.Village, error) {
	return u.repo.FindVillageByID(ctx, id)
}
