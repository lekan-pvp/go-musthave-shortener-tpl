package controllers

import (
	"context"
	"encoding/json"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/cookie_handler"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

func (controller *Controller) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var uuid string
	var in []string
	wg := &sync.WaitGroup{}
	errCh := make(chan error)
	ctx, cancel := context.WithCancel(r.Context())


	cookie, err := r.Cookie("token")
	if err != nil || !cookie_handler.CheckCookie(cookie) {
		cookie = cookie_handler.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	if len(values) != 2 {
		log.Panicln("cookie format error...")
		http.Error(w, err.Error(), 500)
		return
	}
	uuid = values[0]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("reading body error...")
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(body, &in)
	if err != nil {
		log.Println("decoding json error...")
		http.Error(w, err.Error(), 500)
		return
	}

	for _, short := range in {
		wg.Add(1)
		go controller.DeleteURLs(ctx, uuid, short, errCh, wg)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	if err := <-errCh; err != nil {
		log.Println(err)
		cancel()
		return
	}

	w.WriteHeader(202)
}
