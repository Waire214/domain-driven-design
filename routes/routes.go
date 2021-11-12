package routes

import (
	"net/http"
	"presentation/interfaces"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(appPort, hostAddress string, user interfaces.CountryInterface) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	router.Mount("/country", authEndpoint(user))
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(hostAddress+appPort+"/swagger/doc.json"),
	))

	return router
}

func authEndpoint(country interfaces.CountryInterface) http.Handler {
	router := chi.NewRouter()
	router.Post("/create", country.CreateCountry)
	return router
}
