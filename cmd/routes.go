package main

import (
	"net/http"

	"github.com/go-chi/chi"

	piihandler "github.com/firdasafridi/stat-hybrid/internal/handler/pii"
	"github.com/firdasafridi/stat-hybrid/lib/common/log"
	"github.com/firdasafridi/stat-hybrid/lib/common/writer"

)

type moduleHandler struct {
	PIIHandler piihandler.PIIHandler
}

func newRoutes(mHandler moduleHandler) *chi.Mux {

	log.Println("Starting to create new routing...")
	router := chi.NewRouter()

	router.Get("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer.WriteOK(r.Context(), w, "OK")
	}))

	router.Get("/pii", mHandler.PIIHandler.GetPIIData)
	return router
}
