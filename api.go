package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAdress string
}

func NewAPIServer(listenAdress string) *APIServer {
	return &APIServer{
		listenAdress: listenAdress,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/movie", makeHTTPHandleFunc(s.handleMovie))
	router.HandleFunc("/movie/{id}", makeHTTPHandleFunc(s.handleGetMovie))

	log.Println("JSON API server running on port: ", s.listenAdress)
	http.ListenAndServe(s.listenAdress, router)
}

func (s *APIServer) handleMovie(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetMovie(w, r)
	}
	if r.Method == "POST" {
		return s.handleAddMovie(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteMovie(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetMovie(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	return WriteJSON(w, http.StatusOK, &Movie{})
}

func (s *APIServer) handleAddMovie(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteMovie(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleRateMovie(w http.ResponseWriter, r *http.Request) error {
	return nil
}
