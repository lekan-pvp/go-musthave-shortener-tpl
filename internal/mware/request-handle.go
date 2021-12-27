package mware

import (
	"compress/gzip"
	"io"
	"net/http"
)

func RequestHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reader io.Reader

		if r.Header.Get("Content-Encoding") == "gzip" {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error()+"!", http.StatusInternalServerError)
				return
			}
			reader = gz
			defer gz.Close()
		} else {
			reader = r.Body
		}

		req, err := http.NewRequest(r.Method, r.RequestURI, reader)
		if err != nil {
			http.Error(w, err.Error()+"!!", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, req)
	})
}
