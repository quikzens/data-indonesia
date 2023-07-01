package config

import (
	"api/handler"
	"api/repository"
	"api/usecase"
	"net/http"
)

func InitHttpHandler(addr string) *http.Server {
	db := InitGormDatabase()
	r := repository.NewRepository(db)
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)
	router := handler.NewRouter(h)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
