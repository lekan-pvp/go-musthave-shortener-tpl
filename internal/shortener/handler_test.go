package shortener

import (
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handler_GetURLByIDHandler(t *testing.T) {
	type fields struct {
		store *Store
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
				store: &Store{
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
				store: &Store{
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
			router.Route("/{articleID}", func(r chi.Router) {
				r.Get("/", h.GetURLByIDHandler)
			})

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
		store *Store
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
			request: "/",
			fields: fields{
				store: &Store{
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
			defer result.Body.Close()

			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

		})
	}
}
