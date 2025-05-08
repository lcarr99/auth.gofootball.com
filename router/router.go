package router

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	response, error := json.Marshal("json")

	if error != nil {
		return
	}

	http.ResponseWriter.Write(w, response)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/authenticate", authenticateHandler)

	return r
}
