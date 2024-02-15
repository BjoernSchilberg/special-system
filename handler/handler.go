// handler/handler.go
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User struct definieren
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Eine einfache In-Memory-Datenbank für Demonstrationszwecke
var users = []User{
	{ID: "1", Name: "John Doe"},
	{ID: "2", Name: "Jane Doe"},
}

// GetUsers gibt alle Benutzer zurück
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UserByID gibt einen Benutzer anhand der ID zurück
func UserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

// CreateUser fügt einen neuen Benutzer hinzu
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, user) // Einfaches Beispiel, keine ID-Überprüfung
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser entfernt einen Benutzer anhand der ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for i, user := range users {
		if user.ID == id {
			// Benutzer aus der Liste entfernen
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "User with ID %s deleted", id)
			return
		}
	}

	// Wenn kein Benutzer mit der ID gefunden wurde
	http.NotFound(w, r)
}
