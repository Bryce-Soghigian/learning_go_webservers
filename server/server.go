package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message":"put yeet"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"delete yeet"}`))
	case "PATCH":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "patches === testing data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitcks?"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"not found!"}`))
	}

}
func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
