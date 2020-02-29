package main

import (
	"net/http"

	"github.com/er230059/golang-practice/router"
)

func main() {
	router := router.InitRouter()

	s := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	s.ListenAndServe()
}
