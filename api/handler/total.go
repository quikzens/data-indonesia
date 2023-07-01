package handler

import "net/http"

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
