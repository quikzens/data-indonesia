package main

import (
	"net/http"
)

func InitHttpHandler(addr string) *http.Server {
	db := InitGormDatabase()
	r := NewRepository(db)
	u := NewUsecase(r)
	h := NewHandler(u)
	router := NewRouter(h)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
