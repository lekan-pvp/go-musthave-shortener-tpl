package shortener

import (
	"encoding/json"
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
	apiShorten = "/api/shorten"
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
		r.Get("/", h.GetURLByIDHandler)
	})
	router.Post(apiShorten, h.ApiShortenHandler)
}


// GetURLByIDHandler -- возвращает длинный URL из локального хранилища по ключу, которым является короткий URL
func (h *handler) GetURLByIDHandler(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	key := prefix +articleID

	longURL, ok := h.store.Get(key)
	if !ok {
		http.Error(w, "Wrong id", 400)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// GetURLByIDHandler -- создает короткий URL и сохраняет в локальном хранилище
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

type ApiShortenBody struct {
	URL string `json:"-"`
	Result string `json:"result,omitempty"`
}

// ApiShortenHandler -- принимает и возвращает объекты JSON в теле запроса и ответа
func (h *handler) ApiShortenHandler(w http.ResponseWriter, r *http.Request)  {
	var v ApiShortenBody
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	short := h.store.Put(v.URL)

	v.Result = short

	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}
