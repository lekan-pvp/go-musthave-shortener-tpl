package controllers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/testhelper"
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

	res, body := testhelper.TestRequest(t, ts, "POST", "/", nil)
	defer res.Body.Close()
	if string(body) != "http://localhost:8080/whTHc" {
		t.Fatalf("want %s, got %s", "http://localhost:8080/whTHc", string(body))
	}

	res, _ = testhelper.TestRequest(t, ts, "POST", "/", nil)
	defer res.Body.Close()
	if res.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Fatalf("want %s, got %s", "text/plain; charset=utf-8", res.Header.Get("Content-Type"))
	}

	res, _ = testhelper.TestRequest(t, ts, "POST", "/", nil)
	defer res.Body.Close()
	if res.StatusCode != 201 {
		t.Fatalf("want %d, got %d", 201, res.StatusCode)
	}

	res, _ = testhelper.TestRequest(t, ts, "POST", "/somewrong", nil)
	defer res.Body.Close()
	if res.StatusCode != 404 {
		t.Fatalf("want %d, got %d", 404, res.StatusCode)
	}
}
