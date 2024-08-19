package web

import (
	"fmt"
	"net/http"
)

// (GET /)
func GetHome(w http.ResponseWriter, r *http.Request) {
	resp := "Welcome to the home page"

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, resp)
}
