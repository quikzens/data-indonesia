package usecase

import (
	"api/entity"
	"context"
)

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
