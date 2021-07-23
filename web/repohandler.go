package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mitchdennett/myprdroid"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type RepoViewData struct {
	Repos []*myprdroid.Repo
}

func Repos(env *Env, w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	lp := filepath.Join("templates", "base.html")
	fp := filepath.Join("templates", "repos.html")

	tmpl, _ := template.ParseFiles(lp, fp)
	repos, _ := env.RepoService.Repos(os.Getenv("GITACCESSTOKEN"))

	tmpl.ExecuteTemplate(w, "base", RepoViewData{
		Repos: repos,
	})
	return nil
}
