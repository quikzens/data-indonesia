package handler

import (
	"api/repository"
	"api/usecase"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestServer struct {
	*httptest.Server
}

func (ts *TestServer) Get(t *testing.T, urlPath string) (int, http.Header, string) {
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

func NewTestServer(t *testing.T) *TestServer {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	r := repository.NewRepository(db)
	u := usecase.NewUsecase(r)
	h := NewHandler(u)
	router := NewRouter(h)
	ts := httptest.NewServer(router)
	return &TestServer{ts}
}
