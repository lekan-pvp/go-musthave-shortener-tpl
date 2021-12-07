package shortener

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/shortener/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/shortener/storage"
	"io"
	"log"
	"net/http"
)

type handler struct {
	store   *storage.URLStore
	long     BodyRequest
	short  BodyResponse
	baseURL string
}

func NewHandler(cfg *config.Config) *handler {
	store := storage.NewStore(cfg.FileStoragePath)
	log.Println("NewHandler():", cfg.BaseURL)
	return &handler{
		store:   store,
		baseURL: cfg.BaseURL,
	}
}

func (h *handler) Register(router chi.Router) {
	router.Post("/", h.AddShortURLHandler)
	router.Get("/{articleID}", h.GetURLByIDHandler)
	router.Post("/api/shorten", h.APIShortenHandler)

}

// GetURLByIDHandler -- возвращает длинный URL из локального хранилища по ключу, которым является короткий URL
func (h *handler) GetURLByIDHandler(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "articleID")
	url, err := h.store.Get(key)
	log.Println(url)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if url == "" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// AddShortURLHandler -- создает короткий URL и сохраняет в локальном хранилище
func (h *handler) AddShortURLHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	url := string(body)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	key := h.store.Put(url)

	w.Write([]byte(h.baseURL + "/" + key))
}

// APIShortenHandler -- принимает и возвращает объекты JSON в теле запроса и ответа
func (h *handler) APIShortenHandler(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(&h.long); err != nil {
		http.Error(w, err.Error()+"!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	key := h.store.Put(h.long.LongURL)

	h.short.ShortURL = h.baseURL + "/" + key

	result, err := json.Marshal(h.short)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(result))
}
