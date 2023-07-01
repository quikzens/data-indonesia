package handler

import (
	"api/entity"
	"api/helper"
	"net/http"
)

func (h *Handler) GetProvinces(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()

	param := entity.PaginateProvincesParam{
		BasePaginateParam: entity.BasePaginateParam{
			Keywords: queryParams.Get("keywords"),
			SortBy:   queryParams.Get("sort_by"),
			OrderBy:  queryParams.Get("order_by"),
		},
	}

	paginateResult, err := h.usecase.GetProvinces(ctx, param)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, paginateResult.Data, MetaResponse{
		HTTPCode: http.StatusOK,
		Total:    paginateResult.Total,
	})
}

func (h *Handler) PaginateProvinces(w http.ResponseWriter, r *http.Request) {
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

	param := entity.PaginateProvincesParam{
		BasePaginateParam: entity.BasePaginateParam{
			Keywords: queryParams.Get("keywords"),
			SortBy:   queryParams.Get("sort_by"),
			OrderBy:  queryParams.Get("order_by"),
			Limit:    limit,
			Offset:   offset,
		},
	}

	paginateResult, err := h.usecase.PaginateProvinces(ctx, param)
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

func (h *Handler) GetProvince(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := helper.GetUrlIntParam(r, "ID", "ID Must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	province, err := h.usecase.GetProvince(ctx, id)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, province, MetaResponse{
		HTTPCode: http.StatusOK,
	})
}
