package apibin

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
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
	_ = render.Render(w, r, newResponse())
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, newResponse())
}

func patchHandler(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, newResponse())
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, newResponse())
}

type response struct {
	Arguments map[string]string      `json:"args,omitempty"`
	Data      string                 `json:"data,omitempty"`
	Files     map[string][]*file     `json:"files,omitempty"`
	Form      map[string]string      `json:"form,omitempty"`
	Headers   map[string]string      `json:"headers,omitempty"`
	Origin    string                 `json:"origin,omitempty"`
	Json      map[string]interface{} `json:"json,omitempty"`
	Url       string                 `json:"url,omitempty"`
}

type file struct {
	multipart.FileHeader
	Content string
}

func newResponse() *response {
	return &response{
		Arguments: make(map[string]string),
		Files:     make(map[string][]*file),
		Form:      make(map[string]string),
		Headers:   make(map[string]string),
		Json:      make(map[string]interface{}),
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

	if r.Method != http.MethodGet {
		if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			rawData, _ := ioutil.ReadAll(r.Body)
			res.Data = string(rawData)

			_ = json.Unmarshal(rawData, &res.Json)
		}

		if strings.Contains(r.Header.Get("Content-Type"), "application/xml") {
			rawData, _ := ioutil.ReadAll(r.Body)
			res.Data = string(rawData)
		}

		if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			_ = r.ParseMultipartForm(10 * 1024 * 1024)

			for key := range r.Form {
				res.Form[key] = r.Form.Get(key)
			}

			for key := range r.MultipartForm.File {
				for _, f := range r.MultipartForm.File[key] {
					handler, _ := f.Open()
					bytes, _ := ioutil.ReadAll(handler)

					res.Files[key] = append(res.Files[key], &file{
						FileHeader: *f,
						Content: fmt.Sprintf(
							"data:%s;base64,%s",
							f.Header.Get("Content-Type"), base64.StdEncoding.EncodeToString(bytes),
						),
					})
				}
			}
		}
	}

	return nil
}
