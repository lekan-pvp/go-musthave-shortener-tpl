package shortener

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handler_GetURLByIDHandler(t *testing.T) {
	type fields struct {
		store *MemoryStore
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		contentType string
		statusCode int
	}
	tests := []struct {
		name   string
		request string
		fields fields
		args   args
		want want
	}{
		{
			name: "success test",
			request: "http://localhost:8080/XVlBz",
			fields: fields{
				store: &MemoryStore{
					db: map[string]string{"http://localhost:8080/XVlBz": "http://google.com"},
				},
			},
			want: want{
				contentType: "text/plain",
				statusCode: 307,
			},
		},
		{
			name: "not found",
			request: "http://localhost:8080/1",
			fields: fields{
				store: &MemoryStore{
					db: map[string]string{"http://localhost:8080/XVlBz": "http://google.com"},
				},
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				store: tt.fields.store,
			}
			router := chi.NewRouter()
			router.Get("/{articleID}", h.GetURLByIDHandler)

			req := httptest.NewRequest(http.MethodGet, tt.request, nil)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			result := rr.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

		})
	}
}

func Test_handler_CreateShortURLHandler(t *testing.T) {
	type fields struct {
		store *MemoryStore
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		contentType string
		statusCode int
		body string
	}
	tests := []struct {
		name   string
		request string
		fields fields
		args   args
		want want
	}{
		{
			name: "success test",
			request: "/",
			fields: fields{
				store: &MemoryStore{
					db: map[string]string{"":""},
				},
			},
			want: want{
				contentType: "text/plain",
				statusCode: 201,
			},
		},
		{
			name: "wrong endpoint",
			request: "/wrong",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode: 404,
			},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				store: tt.fields.store,
			}

			router := chi.NewRouter()
			router.Post("/", h.CreateShortURLHandler)
			req := httptest.NewRequest(http.MethodPost, tt.request, nil)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			result := rr.Result()

			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

			assert.NotEmpty(t, result.Body)

		})
	}
}

func Test_handler_APIShortenHandler(t *testing.T) {
	type fields struct {
		store  *MemoryStore
		url    BodyRequest
		result BodyResponse
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		contentType string
		statusCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want want
	}{
		{
			name: "success test",
			fields: fields{
				store: &MemoryStore{
					db: map[string]string{"":""},
				},
				url: BodyRequest{
					GoalURL: "http://google.com",
				},
				result: BodyResponse{
					ResultURL: "http://localhost:8080/gbaiC",
				},
			},
			want: want{
				contentType: "application/json; charset=UTF-8",
				statusCode: 201,
			},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				store:  tt.fields.store,
				url:    tt.fields.url,
				result: tt.fields.result,
			}

			router := chi.NewRouter()
			router.Post("/api/shorten", h.APIShortenHandler)

			txBz, err := json.Marshal(tt.fields.url)
			if err != nil {
				log.Fatal(err)
				return
			}

			req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer([]byte(txBz)))
			req.Header.Set("Content-Type", "application/json; charset=UTF-8")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			result := rr.Result()

			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.NotEmpty(t, result.Body)
		})
	}
}
