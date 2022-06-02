package main

import (
	"log"
	"net/http"
	"api/web"
	"api/postgres"
)

func main()  {
	store, err := postgres.NewStore("postgres://postgres:root@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":5010", h)
}