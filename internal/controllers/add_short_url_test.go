package controllers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/test_helper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURLsController_AddURL(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		w.Write([]byte("http://localhost:8080/whTHc"))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	if _, body := test_helper.TestRequest(t, ts, "POST", "/", nil); string(body) != "http://localhost:8080/whTHc" {
		t.Fatalf("want %s, got %s", "http://localhost:8080/whTHc", string(body))
	}

	if res, _ := test_helper.TestRequest(t, ts, "POST", "/", nil); res.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Fatalf("want %s, got %s", "text/plain; charset=utf-8", res.Header.Get("Content-Type"))
	}

	if res, _ := test_helper.TestRequest(t, ts, "POST", "/", nil); res.StatusCode != 201 {
		t.Fatalf("want %d, got %d", 201, res.StatusCode)
	}

	if res, _ := test_helper.TestRequest(t, ts, "POST", "/somewrong", nil); res.StatusCode != 404 {
		t.Fatalf("want %d, got %d", 404, res.StatusCode)
	}
}
