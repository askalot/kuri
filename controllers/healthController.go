package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
)

func HealthShow(w http.ResponseWriter, r *http.Request) {
	MimeTypeApplicationJSON := "application/json"
	healthStatus := struct {
		HealthStatus string `json:"status"`
	}{
		HealthStatus: "OK",
	}

	if strings.Contains(r.Header.Get("Accept"), MimeTypeApplicationJSON) {
		w.Header().Set("Content-Type", MimeTypeApplicationJSON)
		json.NewEncoder(w).Encode(healthStatus)
	} else {
		tmpl := template.Must(template.ParseFiles("views/health/show.html"))

		tmpl.Execute(w, healthStatus)
	}
}
