package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
	"time"
)

var startTime = time.Now()

func HealthShow(w http.ResponseWriter, r *http.Request) {

	MimeTypeApplicationJSON := "application/json"
	healthStatus := struct {
		HealthStatus string `json:"status"`
		Time         string `json:"time"`
		Uptime       string `json:"uptime"`
	}{
		HealthStatus: "ok",
		Time:         time.Now().Format("2006-01-02 15:04:05 MST"),
		Uptime:       time.Since(startTime).Round(time.Second).String(),
	}

	if strings.Contains(r.Header.Get("Accept"), MimeTypeApplicationJSON) {
		w.Header().Set("Content-Type", MimeTypeApplicationJSON)
		json.NewEncoder(w).Encode(healthStatus)
	} else {
		tmpl := template.Must(template.ParseFiles("views/health/show.html"))

		tmpl.Execute(w, healthStatus)
	}
}
