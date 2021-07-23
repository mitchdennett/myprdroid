package web

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchdennett/myprdroid"
	"net/http"
)

type key int
const psKey key = 0

type Env struct {
	RepoService myprdroid.RepoService
}

type Router struct {
	*httprouter.Router
}

func NewRouter() *Router {
	return &Router{httprouter.New()}
}

func (r *Router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

func (r *Router) Post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}

func (r *Router) Put(path string, handler http.Handler) {
	r.PUT(path, wrapHandler(handler))
}

func (r *Router) Delete(path string, handler http.Handler) {
	r.DELETE(path, wrapHandler(handler))
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctxWithParams := context.WithValue(request.Context(), psKey, params)
		rWithPS := request.WithContext(ctxWithParams)
		h.ServeHTTP(writer, rWithPS)
	}
}

type Handler struct {
	*Env
	Handle func(env *Env, w http.ResponseWriter, r *http.Request, ps httprouter.Params) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Handle(h.Env, w, r, nil)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}