package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"messagebird/pkg/model"

	"github.com/gorilla/mux"
)

// Server - contains a method to bootstrap a http server
type Server struct {
}

// NewServer return a pointer to an HTTP Server
func NewServer() *Server {
	return &Server{}
}

// Start - starts the http server
func (server *Server) Start(port string, ch chan error) {
	r := mux.NewRouter()

	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/sms", server.newSMSHandler).Methods(http.MethodPost)

	log.Printf("HTTP server listening on port : %s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)

	ch <- err
}

// newSMSHandler handles POST requests to /api/sms
func (server *Server) newSMSHandler(w http.ResponseWriter, r *http.Request) {
	var sms model.SMS

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&sms)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if valid, err := sms.IsValid(); !valid {
		resp := response{Message: err.Error()}
		json.NewEncoder(w).Encode(resp)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
	}
}
