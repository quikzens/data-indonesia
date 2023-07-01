package handler

import "net/http"

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}
