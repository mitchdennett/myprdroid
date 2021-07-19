package handler

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//UserTokenKey is used to access the user from context
type key int

const psKey key = 0

// Router We could also put *httprouter.Router in a field to not get access to the original methods (GET, POST, etc. in uppercase)
type Router struct {
	*httprouter.Router
}

//Get allows us to wrap all func calls
func (r *Router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

//Post allows us to wrap all func calls
func (r *Router) Post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}

//Put allows us to wrap all func calls
func (r *Router) Put(path string, handler http.Handler) {
	r.PUT(path, wrapHandler(handler))
}

//Delete allows us to wrap all func calls
func (r *Router) Delete(path string, handler http.Handler) {
	r.DELETE(path, wrapHandler(handler))
}

//NewRouter creates a new wrapped Router
func NewRouter() *Router {
	return &Router{httprouter.New()}
}

//WrapHandler is a little bit of Glue to use with HTTPROUTER
func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctxWithParams := context.WithValue(r.Context(), psKey, ps)
		rWithPS := r.WithContext(ctxWithParams)
		h.ServeHTTP(w, rWithPS)
	}
}

type Handle func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	Handle func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ps := r.Context().Value(psKey).(httprouter.Params)
	err := h.Handle(w, r, ps)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}