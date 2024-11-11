package main

import (
	"encoding/json"
	"log"
	"net/http"
	"purdueapi/purdue_api"
	"time"
)

func handleGetDining(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	location := r.URL.Query().Get("location")
	dateStr := r.URL.Query().Get("date")

	// Parse the date (you may need to adjust date format)
	date, err := time.Parse("2006-01-02", dateStr) // Adjust format if needed
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Call GetDining with parsed parameters
	diningInfo, err := purdue_api.GetDining(location, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the result as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(diningInfo); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func main() {
	// Set up routes or handlers that use functions from purdue_api
	http.HandleFunc("/dining", handleGetDining) // Ensure this matches the function from purdue_api package

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
