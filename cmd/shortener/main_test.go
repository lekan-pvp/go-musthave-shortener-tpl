package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/app"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURLByID(t *testing.T) {
	type want struct {
		contentType string
		statusCode int
	}
	tests := []struct {
		name string
		request string
		id int
		long string
		shorts map[int]URLs
		want want
	}{
		{
			name: "success",
			request: "/1",
			id: 1,
			long: "http://google.com",
			want: want{
				contentType: "text/plain",
				statusCode: 307,
			},
		},
		{
			name: "not found",
			request: "/2",
			long: "http://yandex.ru",
			want: want{
				contentType: "text/plain",
				statusCode: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tt := range tests {
				short := app.Shorting(1)
				shorts[tt.id] = URLs{
					Long: tt.long,
					Short: short,
				}

				router := httprouter.New()
				router.GET("/:id", GetURLByID)

				req := httptest.NewRequest(http.MethodGet, tt.request, nil)

				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				result := rr.Result()
				defer result.Body.Close()
				assert.Equal(t, tt.want.statusCode, result.StatusCode)
			}
		})
	}
}

func TestCreateShortURLHandler(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		request string
		want want
	}{
		{
			name: "success",
			request: "/",
			want: want{
				statusCode: 201,
			},
		},
		{
			name: "wrong endpoint",
			request: "/wrong",
			want: want{
				statusCode: 404,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := httprouter.New()
			router.POST("/", CreateShortURLHandler)

			req := httptest.NewRequest(http.MethodPost, tt.request, nil)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			result := rr.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)

		})
	}
}