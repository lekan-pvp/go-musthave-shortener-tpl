package middleware

import (
	"github.com/go-musthave-shortener-tpl/internal/cookie_handler"
	"net/http"
)

func CookieHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		cookie, err := r.Cookie("uid")
		if err != nil {
			cookie = cookie_handler.CreateCookie()
			http.SetCookie(w, cookie)
		}

		if cookie_handler.CheckCookie(cookie) {
			cookie = cookie_handler.CreateCookie()
			http.SetCookie(w, cookie)
		}
		next.ServeHTTP(w, r)

	})
}
