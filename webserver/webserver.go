package webserver

import (
	"encoding/json"
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./webserver/index.html")
	data := map[string]interface{}{
		"message": "Hola desde el servidor JSON",
		"status":  "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		http.Error(w, "Error al codificar JSON", http.StatusInternalServerError)
	}
}
