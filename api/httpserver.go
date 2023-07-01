package main

import (
	"api/repository"
	"api/usecase"
	"net/http"
)

func InitHttpHandler(addr string) *http.Server {
	db := InitGormDatabase()
	r := repository.NewRepository(db)
	u := usecase.NewUsecase(r)
	h := NewHandler(u)
	router := NewRouter(h)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
