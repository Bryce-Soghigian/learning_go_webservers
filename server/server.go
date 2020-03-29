package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message":"get called"}`))
// 	case "POST":
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(`{"message":"post called"}`))
// 	case "PUT":
// 		w.WriteHeader(http.StatusAccepted)
// 		w.Write([]byte(`{"message":"put yeet"}`))
// 	case "DELETE":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message":"delete yeet"}`))
// 	case "PATCH":
// 		w.WriteHeader(http.StatusAccepted)
// 		w.Write([]byte(`{"message": "patches === testing data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitesting data type limitcks?"}`))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`{"message":"not found!"}`))
// 	}

// }
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	port := ":8080"
	// fmt.Println("========Server Listening On PORT" + port + "=========")
	log.Fatal(http.ListenAndServe(port, r))
}
