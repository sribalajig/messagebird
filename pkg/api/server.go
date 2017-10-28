package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"messagebird/pkg/model"
	"messagebird/pkg/service"

	"github.com/gorilla/mux"
)

// Server - contains a method to bootstrap a http server
type Server struct {
	smsService *service.SMSService
}

// NewServer return a pointer to an HTTP Server
func NewServer(smsService *service.SMSService) *Server {
	return &Server{
		smsService: smsService,
	}
}

// Start - starts the http server
func (server *Server) Start(port string, ch chan error) {
	r := mux.NewRouter()

	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/sms", server.newSMSHandler).Methods(http.MethodPost)
	s.HandleFunc("/sms", server.getByRefID).Methods(http.MethodGet)

	log.Printf("HTTP server listening on port : '%s'\n", port)

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
		w.WriteHeader(http.StatusUnprocessableEntity)
		write(response{Message: err.Error()}, w)

		return
	}

	ref := server.smsService.Send(sms)
	w.WriteHeader(http.StatusOK)
	write(response{Message: "SMS Scheduled", Data: referece{ReferenceID: ref}}, w)
}

func (server *Server) getByRefID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var p []string
	var ok bool
	if p, ok = params["refID"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s := server.smsService.Get(p[0])

	if s == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func write(resp response, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(resp)

	w.Header().Set("Content-Type", "application/json")
}
