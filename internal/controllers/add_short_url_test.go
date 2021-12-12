package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-musthave-shortener-tpl/internal/test_helper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURLsController_AddURL(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		w.Write([]byte("http://localhost:8080/whTHc"))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	if _, body := test_helper.TestRequest(t, ts, "POST", "/", nil); body != "http://localhost:8080/whTHc" {
		t.Fatal(body)
	}

	if res, body := test_helper.TestRequest(t, ts, "POST", "/", nil); res.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Fatal(body)
	}

	if res, body := test_helper.TestRequest(t, ts, "POST", "/", nil); res.StatusCode != 201 {
		t.Fatal(body)
	}
}

