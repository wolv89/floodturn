package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func LoadRoutes(mux *http.ServeMux) {

	apiRouter := http.NewServeMux()

	apiRouter.HandleFunc("GET /daterange", handleGetDateRange)

	mux.Handle("/api/", http.StripPrefix("/api", apiRouter))

}

func jsonResponse(w http.ResponseWriter, status int, data any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	output, err := json.Marshal(data)
	if err != nil {
		log.Printf("error marshaling JSON %s", err)
		return
	}

	w.Write(output)

}
