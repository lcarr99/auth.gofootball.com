package main

import (
	"net/http"

	"gofootball.com/router"
)

func main() {
	r := router.NewRouter()

	http.Handle("/", r)
	http.ListenAndServe(":80", r)
}
