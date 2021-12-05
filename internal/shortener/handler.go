package shortener

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/lekan-pvp/go-musthave-shortener-tpl/internal/config"
	"io"
	"log"
	"net/http"
)



type handler struct {
	store *MemoryStore
	url BodyRequest
	result BodyResponse
	baseURL string
}

func NewHandler(cfg *config.Config) *handler {
	store := NewStore()
	log.Println("NewHandler():", cfg.BaseURL)
	return &handler{
		store: store,
		baseURL: cfg.BaseURL,
	}
}

func (h *handler) Register(router chi.Router) {
	router.Post("/", h.CreateShortURLHandler)
	router.Get("/{articleID}", h.GetURLByIDHandler)
	router.Post("/api/shorten", h.APIShortenHandler)

}


// GetURLByIDHandler -- возвращает длинный URL из локального хранилища по ключу, которым является короткий URL
func (h *handler) GetURLByIDHandler(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	key := h.baseURL + "/" + articleID
	log.Println("from handler baseURL: ", h.baseURL)
	longURL, err := h.store.Get(key)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// CreateShortURLHandler -- создает короткий URL и сохраняет в локальном хранилище
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

	short := Shorting(h.baseURL)
	h.store.Put(long, short)

	w.Write([]byte(short))
}

// APIShortenHandler -- принимает и возвращает объекты JSON в теле запроса и ответа
func (h *handler) APIShortenHandler(w http.ResponseWriter, r *http.Request)  {

	if err := json.NewDecoder(r.Body).Decode(&h.url); err != nil {
		http.Error(w, err.Error()+"!", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	short := Shorting(h.baseURL)
	h.store.Put(h.url.GoalURL, short)

	h.result.ResultURL = short

	result, err := json.Marshal(h.result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(result))
}
