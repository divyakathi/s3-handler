package main

import (
	"net/http"
	"s3-file-manager/handlers"
	log "s3-file-manager/logging"
	"s3-file-manager/promutil"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	promutil.InitPrometheus()
	r := mux.NewRouter().StrictSlash(true)
	r.Use(promutil.PrometheusMiddleware)
	r.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/health", handlers.HealthCheck)
	r.HandleFunc("/restart", handlers.RestartApp)

	r.HandleFunc("/s3", handlers.ProcessFile()).Methods(http.MethodPost, http.MethodGet)

	log.Info(log.LogFields{}, "s3-file-handler is running.")

}
