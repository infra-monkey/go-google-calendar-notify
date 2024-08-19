package api

import (
	"encoding/json"
	"net/http"
)

// optional code omitted

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /api/status)
func (Server) GetApiStatus(w http.ResponseWriter, r *http.Request) {
	resp := Status{
		Status: "OK",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// (GET /api/accounts)
func (Server) GetApiAccounts(w http.ResponseWriter, r *http.Request) {
	resp := Accounts{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// (POST /api/accounts)
func (Server) PostApiAccounts(w http.ResponseWriter, r *http.Request) {
	//resp := Accounts{}

	w.WriteHeader(http.StatusOK)
	//_ = json.NewEncoder(w).Encode(resp)
}
