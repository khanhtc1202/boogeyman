package main

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"

	"github.com/gin-gonic/gin/json"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/khanhtc1202/boogeyman/internal/controller"
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/gateway/repository"
	"github.com/khanhtc1202/boogeyman/internal/gateway/service"
	"github.com/khanhtc1202/boogeyman/pkg/io"
)

func main() {
	restCORS := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r := chi.NewRouter()

	r.Use(restCORS.Handler)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Get("/ping", pingHandle)

	r.Get("/search", searchQueryHandle)

	http.ListenAndServe(":3000", r)
}

func pingHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func searchQueryHandle(w http.ResponseWriter, r *http.Request) {
	searchStrategiesRepo := repository.SearchStrategies()
	searchEnginesRepo := repository.SearchEngines(service.EmptyCollectorList())
	textPresenter := NewWebPresenter(w)

	infoSearchCtl := controller.NewInfoSearch(searchStrategiesRepo, searchEnginesRepo, textPresenter)

	queryString := r.URL.Query().Get("q")
	engine := r.URL.Query().Get("e")
	strategy := r.URL.Query().Get("s")

	err := infoSearchCtl.Search(queryString, engine, strategy)
	if err != nil {
		io.Errorln(err)
	}
}

type resultResponse struct {
	Results []queryResultResponse `json:"results"`
}

func ResultResponse(results *domain.QueryResults) *resultResponse {
	var rs []queryResultResponse
	for _, r := range *results {
		rs = append(rs, QueryResultResponse(r.(*domain.UrlBaseResultItem)))
	}
	return &resultResponse{
		Results: rs,
	}
}

type queryResultResponse struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func QueryResultResponse(r *domain.UrlBaseResultItem) queryResultResponse {
	return queryResultResponse{
		Title:       r.GetTitleString(),
		Url:         r.GetUrl(),
		Description: r.GetDescription(),
	}
}

type WebPresenter struct {
	writer http.ResponseWriter
}

func NewWebPresenter(w http.ResponseWriter) *WebPresenter {
	return &WebPresenter{
		writer: w,
	}
}

func (w *WebPresenter) PrintList(results *domain.QueryResults) error {
	response := ResultResponse(results)
	j, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.writer.Write(j)
	return nil
}
