package handler

import (
	"api/helper"
	"net/http"
	"testing"
)

func TestHealthz(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Close()

	statusCode, _, respBody := ts.Get(t, "/healthz")

	helper.Equal(t, statusCode, http.StatusOK)
	helper.Equal(t, respBody, "ok")
}
