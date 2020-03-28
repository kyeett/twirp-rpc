package main

import (
	"log"
	"net/http"

	"github.com/kyeett/twirp-rpc/common/logger"
	"github.com/kyeett/twirp-rpc/common/router"
	"github.com/kyeett/twirp-rpc/examples/template/internal/templateserver"
	"github.com/kyeett/twirp-rpc/examples/template/rpc/template"
)

func main() {
	l := logger.NewDefault()
	r := router.NewDefault()

	// Server
	server := templateserver.New(l)
	twirpHandler := template.NewTemplateServicerServer(server, nil)
	r.Mount(template.TemplateServicerPathPrefix, twirpHandler)

	l.Info("Starting server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	l.Info("Shutting down server")
}
