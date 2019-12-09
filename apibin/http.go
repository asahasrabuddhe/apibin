package apibin

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func httpMethods() http.Handler {
	r := chi.NewRouter()

	r.Get("/get", getHandler)
	r.Post("/post", postHandler)
	r.Put("/put", putHandler)
	r.Patch("/patch", patchHandler)
	r.Delete("/delete", deleteHandler)

	return r
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	 _ = render.Render(w, r, newResponse())
}

func postHandler(w http.ResponseWriter, r *http.Request) {

}

func putHandler(w http.ResponseWriter, r *http.Request) {

}

func patchHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

}

type response struct {
	Arguments map[string]string `json:"args"`
	Headers   map[string]string `json:"headers"`
	Origin    string            `json:"origin"`
	Url       string            `json:"url"`
}

func newResponse() *response {
	return &response{
		Arguments: make(map[string]string),
		Headers:   make(map[string]string),
	}
}

func (res *response) Bind(r *http.Request) error {
	panic("implement me")
}

func (res *response) Render(w http.ResponseWriter, r *http.Request) error {
	for key := range r.URL.Query() {
		res.Arguments[key] = r.URL.Query().Get(key)
	}

	for key := range r.Header {
		res.Headers[key] = r.Header.Get(key)
	}
	res.Headers["Host"] = r.Host

	res.Origin = r.RemoteAddr
	res.Url = r.RequestURI

	return nil
}
