package web

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func Index(env *Env, w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	lp := filepath.Join("templates", "base.html")
	fp := filepath.Join("templates", "index.html")

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "base", nil)
	return nil
}