package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rohan3011/go-server/services/user"
	"github.com/rohan3011/go-server/services/view"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr,
		db,
	}
}

// Run starts the API server and listens for incoming requests.
//
// It initializes the router, sets up middleware, mounts the API routes,
// sets up the view service, sets up the user service, and starts the server
// to listen on the specified address.
//
// Returns an error if there was a problem starting the server.
func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	apiRouter := chi.NewRouter()
	router.Mount("/api/v1", apiRouter)

	// view service
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))
	viewHandler := view.NewHandler()
	viewHandler.RegisterRoutes(router)

	// user service
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(apiRouter)

	log.Printf("Listening on http://localhost%s", s.addr)
	return http.ListenAndServe(s.addr, router)
}
