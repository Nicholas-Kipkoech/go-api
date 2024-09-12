package main

import (
	"encoding/json"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getClientProfile(w, r)
	case http.MethodPatch:
		updateClientProfile(w, r)
	default:
		http.Error(w, "Forbidden", http.StatusMethodNotAllowed)
	}

}

func getClientProfile(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId")
	clientProfile, ok := database[clientId]

	if !ok || clientId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := ClientProfile{
		Email: clientProfile.Email,
		Name:  clientProfile.Name,
		Id:    clientProfile.Id,
	}
	json.NewEncoder(w).Encode(response)
}

func updateClientProfile(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId")
	clientProfile, ok := database[clientId]
	if !ok || clientId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	var payloadData ClientProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if payloadData.Email != "" {
		clientProfile.Email = payloadData.Email
	}
	if payloadData.Name != "" {
		clientProfile.Name = payloadData.Name
	}
	database[clientProfile.Id] = clientProfile

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clientProfile)

}
