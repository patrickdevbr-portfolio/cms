package rest

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func writeErr(w http.ResponseWriter, err error) error {
	return writeJSON(w, http.StatusInternalServerError, map[string]any{"err": err})
}

func readJSON(w http.ResponseWriter, r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"err": err})
		return err
	}
	defer r.Body.Close()
	return nil
}
