package taktik

import (
	"encoding/json"
	"net/http"
)


func do_index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(Response{
		Ok:     true,
		Status: 200,
	})
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
