package router

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"gofootball.com/database"
)

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	db, dbErr := database.NewDatabase()

	if dbErr != nil {
		w.WriteHeader(500)
	}

	statement, _ := db.Prepare("SELECT uuid FROM files WHERE id = 1")

	var uuid string

	statement.QueryRow().Scan(&uuid)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	response, error := json.Marshal(uuid)

	if error != nil {
		return
	}

	http.ResponseWriter.Write(w, response)
}

func initDatabaseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := database.NewDatabase()

		if err != nil {
			http.Error(w, "Unexpected Error", 500)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(initDatabaseMiddleware)
	r.HandleFunc("/authenticate", authenticateHandler)

	return r
}
