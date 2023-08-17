package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/api/users", withMiddleware(createUser))
	http.HandleFunc("/api/users/", withMiddleware(getUser))

	http.ListenAndServe(":8080", nil)
}

// Custom middleware function
func withMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Perform actions before passing the request to the next handler
		fmt.Println("Middleware: Request received")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Perform actions after the next handler has completed
		fmt.Println("Middleware: Request completed")
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate user data
	if user.Name == "" || user.Email == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}

	// Save the user to the database (not implemented in this example)

	// Simulate successful response for demonstration purposes
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User created successfully"}
	json.NewEncoder(w).Encode(response)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract the value of "id" parameter from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/api/users/")

	// Perform user retrieval based on the provided ID (not implemented in this example)
	if id == "" {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	// Simulate response for demonstration purposes
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"user_id": id}
	json.NewEncoder(w).Encode(response)
}
