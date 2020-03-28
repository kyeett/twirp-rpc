package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kyeett/twirp-rpc/examples/template/internal/templateserver"
	"github.com/kyeett/twirp-rpc/examples/template/rpc/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	server := templateserver.New()
	twirpHandler := template.NewTemplateServicerServer(server, nil)
	routes := chi.NewRouter()
	routes.Use(middleware.RequestID)
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)
	routes.Mount(template.TemplateServicerPathPrefix, twirpHandler)
	fmt.Println("Starting server")
	if err := http.ListenAndServe(":8080", routes); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shutting down server")
}
