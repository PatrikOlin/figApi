package main

import (
	// "fmt"
	"os"
	"time"
	"path/filepath"
	"log"
	"net/http"
	// "strconv"
	// "figApi/util"
	// "figApi/datastore"											  

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type server struct{}

// func get(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "tjenare världen"}`))
// }

// func post(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte(`{"message": "tjenare världen post"}`))
// }

// func put(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusAccepted)
// 	w.Write([]byte(`{"message": "tjenare världen put"}`))
// }

// func delete(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "tjenare världen delete"}`))
// }

// func params(w http.ResponseWriter, r *http.Request) {
// 	pathParams := mux.Vars(r)
// 	w.Header().Set("Content-Type", "application/json")

// 	userID := -1
// 	var err error
// 	if val, ok := pathParams["userID"]; ok {
// 		userID, err = strconv.Atoi(val)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(`{"message": "need a number"}`))
// 			return
// 		}
// 	}

// 	commentID := -1
// 	if val, ok := pathParams["commentID"]; ok {
// 		commentID, err = strconv.Atoi(val)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(`{"message": "need a number"}`))
// 			return
// 		}
// 	}

// 	query := r.URL.Query()
// 	location := query.Get("location")

// 	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
// }

func main() {
	// datastore.Initdb()										  	  
	t := time.Now()
	logpath := filepath.Join(".", "logs")
	os.MkdirAll(logpath, os.ModePerm)
	logFile, err := os.OpenFile(logpath + "/fig_" + t.Format("20060102") + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/articles", getArticle).Methods(http.MethodGet)
	api.HandleFunc("/people", getPerson).Methods(http.MethodGet)
	api.HandleFunc("/companies", getCompany).Methods(http.MethodGet)
	// api.HandleFunc("/", put).Methods(http.MethodPut)
	// api.HandleFunc("/", delete).Methods(http.MethodDelete)

	// api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8124", handlers.CombinedLoggingHandler(logFile, handlers.CORS() (r))) )
}
