package main

import (
	"github.com/mitchdennett/myprdroid/handler"
	"log"
	"net/http"
)

func main() {

	router := handler.NewRouter()

	router.ServeFiles("/static/*filepath", http.Dir("static/public"))

	router.Get("/", handler.Handler{Handle: handler.Index})

	log.Fatal(http.ListenAndServe(":8000", router))
}