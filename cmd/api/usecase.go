package main

import (
	"context"
)

type Usecase struct {
	repo *Repository
}

func NewUsecase(repo *Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetProvinces(ctx context.Context, param PaginateProvincesParam) (PaginateResult[Province], error) {
	provinces, count, err := u.repo.SearchProvinces(ctx, param)
	if err != nil {
		return PaginateResult[Province]{}, err
	}

	return PaginateResult[Province]{
		Data:   provinces,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) PaginateProvinces(ctx context.Context, param PaginateProvincesParam) (PaginateResult[Province], error) {
	if param.Limit == 0 {
		param.Limit = 10 // fallback to default 10 paginate
	}

	provinces, count, err := u.repo.SearchProvinces(ctx, param)
	if err != nil {
		return PaginateResult[Province]{}, err
	}

	return PaginateResult[Province]{
		Data:   provinces,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) GetProvince(ctx context.Context, id uint) (Province, error) {
	return u.repo.FindProvinceByID(ctx, id)
}

func (u *Usecase) GetCities(ctx context.Context, param PaginateCitiesParam) (PaginateResult[City], error) {
	cities, count, err := u.repo.SearchCities(ctx, param)
	if err != nil {
		return PaginateResult[City]{}, err
	}

	return PaginateResult[City]{
		Data:   cities,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) PaginateCities(ctx context.Context, param PaginateCitiesParam) (PaginateResult[City], error) {
	if param.Limit == 0 {
		param.Limit = 10 // fallback to default 10 paginate
	}

	cities, count, err := u.repo.SearchCities(ctx, param)
	if err != nil {
		return PaginateResult[City]{}, err
	}

	return PaginateResult[City]{
		Data:   cities,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) GetCity(ctx context.Context, id uint) (City, error) {
	return u.repo.FindCityByID(ctx, id)
}

func (u *Usecase) GetSubdistricts(ctx context.Context, param PaginateSubdistrictsParam) (PaginateResult[Subdistrict], error) {
	subdistricts, count, err := u.repo.SearchSubdistricts(ctx, param)
	if err != nil {
		return PaginateResult[Subdistrict]{}, err
	}

	return PaginateResult[Subdistrict]{
		Data:   subdistricts,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) PaginateSubdistricts(ctx context.Context, param PaginateSubdistrictsParam) (PaginateResult[Subdistrict], error) {
	if param.Limit == 0 {
		param.Limit = 10 // fallback to default 10 paginate
	}

	subdistricts, count, err := u.repo.SearchSubdistricts(ctx, param)
	if err != nil {
		return PaginateResult[Subdistrict]{}, err
	}

	return PaginateResult[Subdistrict]{
		Data:   subdistricts,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) GetSubdistrict(ctx context.Context, id uint) (Subdistrict, error) {
	return u.repo.FindSubdistrictByID(ctx, id)
}

func (u *Usecase) GetVillages(ctx context.Context, param PaginateVillagesParam) (PaginateResult[Village], error) {
	villages, count, err := u.repo.SearchVillages(ctx, param)
	if err != nil {
		return PaginateResult[Village]{}, err
	}

	return PaginateResult[Village]{
		Data:   villages,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) PaginateVillages(ctx context.Context, param PaginateVillagesParam) (PaginateResult[Village], error) {
	if param.Limit == 0 {
		param.Limit = 10 // fallback to default 10 paginate
	}

	villages, count, err := u.repo.SearchVillages(ctx, param)
	if err != nil {
		return PaginateResult[Village]{}, err
	}

	return PaginateResult[Village]{
		Data:   villages,
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  count,
	}, nil
}

func (u *Usecase) GetVillage(ctx context.Context, id uint) (Village, error) {
	return u.repo.FindVillageByID(ctx, id)
}
