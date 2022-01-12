package controllers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/test_helper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURLsController_GetURLByID1(t *testing.T) {
	store := make(map[string]string)

	store["gbaiC"] = "http://google.com"

	r := chi.NewRouter()
	r.Get("/{articleID}", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "articleID")
		url := store[key]
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(307)
		w.Write([]byte(url))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	if res, _ := test_helper.TestRequest(t, ts, "GET", "/gbaiC", nil); res.StatusCode != 307 {
		t.Fatalf("want %d, got %d", 307, res.StatusCode)
	}

	if res, _ := test_helper.TestRequest(t, ts, "GET", "/user/7", nil); res.StatusCode != 404 {
		t.Fatalf("want %d, got %d", 404, res.StatusCode)
	}

	if res, _ := test_helper.TestRequest(t, ts, "GET", "/gbaiC", nil); res.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Fatalf("want %s, got %s", "text/plain; charset=utf-8", res.Header.Get("Content-Type"))
	}

	if _, body := test_helper.TestRequest(t, ts, "GET", "/gbaiC", nil); string(body) != "http://google.com" {
		t.Fatalf("want %s, got %s", "http://google.com", string(body))
	}
}
