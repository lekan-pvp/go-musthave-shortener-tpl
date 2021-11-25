package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/handlers"
	"io"
	"log"
	"net/http"
)

var _ handlers.Handler = &handler{}

const (
	createURL = "/"
	getURL = "/{articleID}"
)

type handler struct {
	store *Store
}

func NewHandler() handlers.Handler {
	store := NewStore()
	return &handler{store: store}
}

func (h *handler) Register(router chi.Router) {
	router.Post(createURL, h.CreateShortURLHandler)
	router.Route(getURL, func(r chi.Router) {
		r.Get("/", h.GetURLByID)
	})
}


func (h *handler) GetURLByID(w http.ResponseWriter, r *http.Request) {
	articleID := r.URL.Path
	key := prefix[:len(prefix)-1]+articleID
	log.Println(key)
	//
	longURL := h.store.Get(key)
	log.Println(longURL)


	//id, err := strconv.Atoi(param)
	//log.Println(id)
	//if err != nil {
	//	http.Error(w, "Wrong", 400)
	//	return
	//}
	//
	//long := shorts[id].Long
	//log.Println(long)
	//if long == "" {
	//	http.Error(w, "Wrong id", 400)
	//	return
	//}
	//w.Header().Set("Content-Type", "text/plain")
	//w.Header().Set("Location", long)
	//w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte("get"))
}

func (h *handler) CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	long := string(body)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	short := h.store.Put(long)
	long = h.store.Get(short)
	log.Printf("%s\n", h.store.db[short])
	w.Write([]byte(short))
}


