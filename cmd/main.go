package main

import (
	"github.com/joho/godotenv"
	"github.com/mitchdennett/myprdroid/github"
	"github.com/mitchdennett/myprdroid/web"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := web.NewRouter()

	env := &web.Env{
		RepoService: &github.RepoService{},
	}

	router.Get("/", web.Handler{Env: env, Handle: web.Repos})

	log.Fatal(http.ListenAndServe(":8000", router))
}
