package pagination

import (
	"context"
	"net/http"
	"strconv"
)

const (
	OptionsContextKey = "pagination_options"
)

func Middleware(h http.HandlerFunc, defaultPtoken int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ptokenFromQuery := r.URL.Query().Get("ptoken")
		ptoken := defaultPtoken
		var ptokenParseErr error
		if ptokenFromQuery != "" {
			if ptoken, ptokenParseErr = strconv.Atoi(ptokenFromQuery); ptokenParseErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("bad token"))
				return
			}

		}

		options1 := POptions{
			Ptoken: ptoken,
		}
		ctx := context.WithValue(r.Context(), OptionsContextKey, options1)
		r1 := r.WithContext(ctx)

		h(w, r1)
	}

}

type POptions struct {
	Ptoken int
}

