package main

import (
	"api/entity"
	"api/repository"
	"context"
)

type Usecase struct {
	repo *repository.Repository
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

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

func (u *Usecase) GetTotals(ctx context.Context) (entity.GetTotalResult, error) {
	var result entity.GetTotalResult
	var err error

	_, result.Provinces, err = u.repo.SearchProvinces(ctx, entity.PaginateProvincesParam{})
	if err != nil {
		return entity.GetTotalResult{}, err
	}

	_, result.Cities, err = u.repo.SearchCities(ctx, entity.PaginateCitiesParam{})
	if err != nil {
		return entity.GetTotalResult{}, err
	}

	_, result.Subdistricts, err = u.repo.SearchSubdistricts(ctx, entity.PaginateSubdistrictsParam{})
	if err != nil {
		return entity.GetTotalResult{}, err
	}

	_, result.Villages, err = u.repo.SearchVillages(ctx, entity.PaginateVillagesParam{})
	if err != nil {
		return entity.GetTotalResult{}, err
	}

	return result, nil
}
