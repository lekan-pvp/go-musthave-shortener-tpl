package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/testHelper"
	"net/http"
	"net/http/httptest"
	"testing"
)

type shorttest struct {
	Key string `json:"result"`
}

type bodytest struct {
	Result string `json:"result"`
}

func TestURLsController_APIShorten(t *testing.T) {
	short := shorttest{
		Key: "http://localhost:8080/gbaiC",
	}

	r := chi.NewRouter()
	r.Post("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(201)
		result, err := json.Marshal(&short)
		if err != nil {
			t.Error(err.Error())
			return
		}

		w.Write([]byte(result))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, _ := testHelper.TestRequest(t, ts, "POST", "/api/shorten", nil)
	defer res.Body.Close()
	if res.StatusCode != 201 {
		t.Fatalf("want %d, got %d\n", 201, res.StatusCode)
	}

	res, _ = testHelper.TestRequest(t, ts, "POST", "/api/shorten", nil)
	defer res.Body.Close()
	if res.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Fatalf("want %s, got %s", "application/json; charset=utf-8", res.Header.Get("Content-Type"))
	}

	res, _ = testHelper.TestRequest(t, ts, "POST", "/api/shorten/3", nil)
	defer res.Body.Close()
	if res.StatusCode != 404 {
		t.Fatalf("want %d, got %d", 404, res.StatusCode)
	}

	bodytest := bodytest{}

	_, body := testHelper.TestRequest(t, ts, "POST", "/api/shorten", nil)

	err := json.Unmarshal(body, &bodytest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	bodyRes := bodytest.Result
	if string(bodyRes) != "http://localhost:8080/gbaiC" {
		t.Fatalf("want %s, got %s", "http://localhost:8080/gbaiC", bodyRes)
	}
}
