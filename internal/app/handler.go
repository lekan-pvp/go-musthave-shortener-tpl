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
	articleID := chi.URLParam(r, "articleID")
	key := prefix+articleID

	longURL := h.store.Get(key)

	if longURL == "" {
		http.Error(w, "Wrong id", 400)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
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


	log.Printf("%v: %v\n", short, h.store.db[short])
	w.Write([]byte(short))
}


