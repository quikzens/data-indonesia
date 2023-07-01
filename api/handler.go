package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	usecase *Usecase
}

func NewHandler(usecase *Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type MetaResponse struct {
	HTTPCode int `json:"http_code"`
	Limit    int `json:"limit,omitempty"`
	Offset   int `json:"offset,omitempty"`
	Total    int `json:"total,omitempty"`
}

type Response struct {
	Message string              `json:"message,omitempty"`
	Data    interface{}         `json:"data,omitempty"`
	Errors  []HttpResponseError `json:"errors,omitempty"`
	Meta    MetaResponse        `json:"meta"`
}

func (h *Handler) writeSuccess(w http.ResponseWriter, data interface{}, meta MetaResponse) {
	res := Response{
		Data: data,
		Meta: meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(meta.HTTPCode)
	responseBody, _ := json.Marshal(res)
	_, _ = w.Write(responseBody)
}

func (h *Handler) translateError(err error) (int, []HttpResponseError) {
	switch origErr := err.(type) {
	case HttpError:
		return origErr.ToHttpError()
	default:
		return InternalServerError{Message: err.Error()}.ToHttpError()
	}
}

func (h *Handler) writeError(w http.ResponseWriter, err error) {
	statusCode, httpErrors := h.translateError(err)
	meta := MetaResponse{
		HTTPCode: statusCode,
	}
	res := Response{
		Errors: httpErrors,
		Meta: MetaResponse{
			HTTPCode: statusCode,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(meta.HTTPCode)
	responseBody, _ := json.Marshal(res)
	_, _ = w.Write(responseBody)
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func (h *Handler) GetProvinces(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()

	param := PaginateProvincesParam{
		BasePaginateParam: BasePaginateParam{
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

	limit, err := GetQueryIntParam(r, "limit", "limit must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	offset, err := GetQueryIntParam(r, "offset", "offset must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := PaginateProvincesParam{
		BasePaginateParam: BasePaginateParam{
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
	id, err := GetUrlIntParam(r, "ID", "ID Must be a number")
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

func (h *Handler) GetCities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()

	provinceId, err := GetQueryIntParam(r, "province_id", "province_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := PaginateCitiesParam{
		BasePaginateParam: BasePaginateParam{
			Keywords: queryParams.Get("keywords"),
			SortBy:   queryParams.Get("sort_by"),
			OrderBy:  queryParams.Get("order_by"),
		},
		ProvinceId: provinceId,
	}

	paginateResult, err := h.usecase.GetCities(ctx, param)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, paginateResult.Data, MetaResponse{
		HTTPCode: http.StatusOK,
		Total:    paginateResult.Total,
	})
}

func (h *Handler) PaginateCities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()

	limit, err := GetQueryIntParam(r, "limit", "limit must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	offset, err := GetQueryIntParam(r, "offset", "offset must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	provinceId, err := GetQueryIntParam(r, "province_id", "province_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := PaginateCitiesParam{
		BasePaginateParam: BasePaginateParam{
			Keywords: queryParams.Get("keywords"),
			SortBy:   queryParams.Get("sort_by"),
			OrderBy:  queryParams.Get("order_by"),
			Limit:    limit,
			Offset:   offset,
		},
		ProvinceId: provinceId,
	}

	paginateResult, err := h.usecase.PaginateCities(ctx, param)
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

func (h *Handler) GetCity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := GetUrlIntParam(r, "ID", "ID Must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	city, err := h.usecase.GetCity(ctx, id)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, city, MetaResponse{
		HTTPCode: http.StatusOK,
	})
}

func (h *Handler) GetSubdistricts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()

	provinceId, err := GetQueryIntParam(r, "province_id", "province_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	cityId, err := GetQueryIntParam(r, "city_id", "city_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := PaginateSubdistrictsParam{
		BasePaginateParam: BasePaginateParam{
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

	limit, err := GetQueryIntParam(r, "limit", "limit must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	offset, err := GetQueryIntParam(r, "offset", "offset must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	provinceId, err := GetQueryIntParam(r, "province_id", "province_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	cityId, err := GetQueryIntParam(r, "city_id", "city_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := PaginateSubdistrictsParam{
		BasePaginateParam: BasePaginateParam{
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
	id, err := GetUrlIntParam(r, "ID", "ID Must be a number")
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

func (h *Handler) GetVillages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParams := r.URL.Query()

	provinceId, err := GetQueryIntParam(r, "province_id", "province_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	cityId, err := GetQueryIntParam(r, "city_id", "city_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	subdistrictId, err := GetQueryIntParam(r, "subdistrict_id", "subdistrict_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := PaginateVillagesParam{
		BasePaginateParam: BasePaginateParam{
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

	limit, err := GetQueryIntParam(r, "limit", "limit must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	offset, err := GetQueryIntParam(r, "offset", "offset must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	provinceId, err := GetQueryIntParam(r, "province_id", "province_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	cityId, err := GetQueryIntParam(r, "city_id", "city_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	subdistrictId, err := GetQueryIntParam(r, "subdistrict_id", "subdistrict_id must be a number")
	if err != nil {
		h.writeError(w, err)
		return
	}

	param := PaginateVillagesParam{
		BasePaginateParam: BasePaginateParam{
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
	id, err := GetUrlIntParam(r, "ID", "ID Must be a number")
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

func (h *Handler) GetTotals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := h.usecase.GetTotals(ctx)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeSuccess(w, result, MetaResponse{
		HTTPCode: http.StatusOK,
	})
}
