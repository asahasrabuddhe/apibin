package apibin

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func statusCodes() http.Handler {
	r := chi.NewRouter()

	r.Get("/{code}", statusCodeHandler)
	r.Post("/{code}", statusCodeHandler)
	r.Put("/{code}", statusCodeHandler)
	r.Patch("/{code}", statusCodeHandler)
	r.Delete("/{code}", statusCodeHandler)

	return r
}

func statusCodeHandler(w http.ResponseWriter, r *http.Request) {
	if code, err := strconv.Atoi(chi.URLParam(r, "code")); err == nil {
		// Only respond for 1xx, 2xx, 3xx, 4xx and 5xx codes
		if code >= 100 && code < 600 {
			w.WriteHeader(code)
		}

		return
	}

	w.WriteHeader(http.StatusBadRequest)
	_, _ = fmt.Fprintf(w, "Invalid status code")
}
