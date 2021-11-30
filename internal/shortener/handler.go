package shortener

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)



type handler struct {
	store *MemoryStore
	url BodyRequest
	result BodyResponse
}

func NewHandler() *handler {
	store := NewStore()
	return &handler{store: store}
}

func (h *handler) Register(router chi.Router) {
	router.Post("/", h.CreateShortURLHandler)
	router.Get("/{articleID}", h.GetURLByIDHandler)
	router.Post("/api/shorten", h.ApiShortenHandler)
}


// GetURLByIDHandler -- возвращает длинный URL из локального хранилища по ключу, которым является короткий URL
func (h *handler) GetURLByIDHandler(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	key := prefix +articleID

	longURL, err := h.store.Get(key)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// GetURLByIDHandler -- создает короткий URL и сохраняет в локальном хранилище
func (h *handler) CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
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

// ApiShortenHandler -- принимает и возвращает объекты JSON в теле запроса и ответа
func (h *handler) ApiShortenHandler(w http.ResponseWriter, r *http.Request)  {

	if err := json.NewDecoder(r.Body).Decode(&h.url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", h.url.URL)
	w.WriteHeader(http.StatusCreated)

	short := h.store.Put(h.url.URL)

	h.result.Result = short

	result, err := json.Marshal(h.result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(result)
}
