package usecase

import (
	"api/entity"
	"context"
)

func (u *Usecase) GetSubdistricts(ctx context.Context, param entity.PaginateSubdistrictsParam) (entity.PaginateResult[entity.Subdistrict], error) {
	subdistricts, count, err := u.repo.SearchSubdistricts(ctx, param)
	if err != nil {
		return entity.PaginateResult[entity.Subdistrict]{}, err
	}

	return entity.PaginateResult[entity.Subdistrict]{
		Data:   subdistricts,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) PaginateSubdistricts(ctx context.Context, param entity.PaginateSubdistrictsParam) (entity.PaginateResult[entity.Subdistrict], error) {
	if param.Limit == 0 {
		param.Limit = 10 // fallback to default 10 paginate
	}

	subdistricts, count, err := u.repo.SearchSubdistricts(ctx, param)
	if err != nil {
		return entity.PaginateResult[entity.Subdistrict]{}, err
	}

	return entity.PaginateResult[entity.Subdistrict]{
		Data:   subdistricts,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) GetSubdistrict(ctx context.Context, id int) (entity.Subdistrict, error) {
	return u.repo.FindSubdistrictByID(ctx, id)
}
