package handler

import (
	"api/entity"
	"api/helper"
	"net/http"
)

func (h *Handler) GetSubdistricts(w http.ResponseWriter, r *http.Request) {
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

	param := entity.PaginateSubdistrictsParam{
		BasePaginateParam: entity.BasePaginateParam{
			Keywords: queryParams.Get("keywords"),
			SortBy:   queryParams.Get("sort_by"),
			OrderBy:  queryParams.Get("order_by"),
		},
		ProvinceId: provinceId,
		CityId:     cityId,
	}

	paginateResult, err := h.usecase.GetSubdistricts(ctx, param)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, paginateResult.Data, MetaResponse{
		HTTPCode: http.StatusOK,
		Total:    paginateResult.Total,
	})
}

func (h *Handler) PaginateSubdistricts(w http.ResponseWriter, r *http.Request) {
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

	param := entity.PaginateSubdistrictsParam{
		BasePaginateParam: entity.BasePaginateParam{
			Keywords: queryParams.Get("keywords"),
			SortBy:   queryParams.Get("sort_by"),
			OrderBy:  queryParams.Get("order_by"),
			Limit:    limit,
			Offset:   offset,
		},
		ProvinceId: provinceId,
		CityId:     cityId,
	}

	paginateResult, err := h.usecase.PaginateSubdistricts(ctx, param)
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

func (h *Handler) GetSubdistrict(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := helper.GetUrlIntParam(r, "ID", "ID Must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	subdistrict, err := h.usecase.GetSubdistrict(ctx, id)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, subdistrict, MetaResponse{
		HTTPCode: http.StatusOK,
	})
}
