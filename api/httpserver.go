package main

import (
	"api/repository"
	"net/http"
)

func InitHttpHandler(addr string) *http.Server {
	db := InitGormDatabase()
	r := repository.NewRepository(db)
	u := NewUsecase(r)
	h := NewHandler(u)
	router := NewRouter(h)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
