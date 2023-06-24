package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

type testServer struct {
	*httptest.Server
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

func newTestServer(t *testing.T) *testServer {
	_ = godotenv.Load()

	db := InitGormDatabase()
	r := NewRepository(db)
	u := NewUsecase(r)
	h := NewHandler(u)
	router := NewRouter(h)

	ts := httptest.NewServer(router)
	return &testServer{ts}
}

func TestHealthz(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	statusCode, _, respBody := ts.get(t, "/healthz")

	Equal(t, statusCode, http.StatusOK)
	Equal(t, respBody, "ok")
}
