package handler

import (
	"api/entity"
	"api/helper"
	"net/http"
)

func (h *Handler) GetVillages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()

	provinceId, err := helper.GetQueryIntParam(r, "province_id", "province_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	cityId, err := helper.GetQueryIntParam(r, "city_id", "city_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	subdistrictId, err := helper.GetQueryIntParam(r, "subdistrict_id", "subdistrict_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := entity.PaginateVillagesParam{
		BasePaginateParam: entity.BasePaginateParam{
			Keywords: queryParams.Get("keywords"),
			SortBy:   queryParams.Get("sort_by"),
			OrderBy:  queryParams.Get("order_by"),
		},
		ProvinceId:    provinceId,
		CityId:        cityId,
		SubdistrictId: subdistrictId,
	}

	paginateResult, err := h.usecase.GetVillages(ctx, param)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, paginateResult.Data, MetaResponse{
		HTTPCode: http.StatusOK,
		Total:    paginateResult.Total,
	})
}

func (h *Handler) PaginateVillages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()

	limit, err := helper.GetQueryIntParam(r, "limit", "limit must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	offset, err := helper.GetQueryIntParam(r, "offset", "offset must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	provinceId, err := helper.GetQueryIntParam(r, "province_id", "province_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	cityId, err := helper.GetQueryIntParam(r, "city_id", "city_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	subdistrictId, err := helper.GetQueryIntParam(r, "subdistrict_id", "subdistrict_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := entity.PaginateVillagesParam{
		BasePaginateParam: entity.BasePaginateParam{
			Keywords: queryParams.Get("keywords"),
			SortBy:   queryParams.Get("sort_by"),
			OrderBy:  queryParams.Get("order_by"),
			Limit:    limit,
			Offset:   offset,
		},
		ProvinceId:    provinceId,
		CityId:        cityId,
		SubdistrictId: subdistrictId,
	}

	paginateResult, err := h.usecase.PaginateVillages(ctx, param)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, paginateResult.Data, MetaResponse{
		HTTPCode: http.StatusOK,
		Limit:    paginateResult.Limit,
		Offset:   paginateResult.Offset,
		Total:    paginateResult.Total,
	})
}

func (h *Handler) GetVillage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := helper.GetUrlIntParam(r, "ID", "ID Must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	village, err := h.usecase.GetVillage(ctx, id)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, village, MetaResponse{
		HTTPCode: http.StatusOK,
	})
}
