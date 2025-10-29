package promutil

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func InitPrometheus() {
	err := prometheus.Register(Endpointcounter)
	err = prometheus.Register(TxnCounter)
	err = prometheus.Register(SuccessfulTxnCounter)
	err = prometheus.Register(TxnLatencyGauge)
	err = prometheus.Register(DataIngestionGauge)
	err = prometheus.Register(UploadCounter)
	err = prometheus.Register(SuccessfulUploadCounter)

	if err != nil {
		fmt.Println("error registering prometheus metric, total requests")
	}
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		var hitPath = r.URL.Path
		if hitPath != "/health" && hitPath != "/restart" && hitPath != "/metrics" {
			Endpointcounter.WithLabelValues(hitPath).Inc()
		}
	})
}

//SLA metrics

// TxnCounter count transactions
var TxnCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "txn_counter",
		Help: "counts all tranactions",
	},
)

// SuccessfulTxnCounter counts successful transactions
var SuccessfulTxnCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "successful_txn_counter",
		Help: "counts all successful tranactions",
	},
)

// UploadCounter count attempts to upload to s3 bucket
var UploadCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "upload_counter",
		Help: "count attempts to upload to s3 bucket",
	},
)

// SuccessfulUploadCounter count attempts to uploads to s3 bucket
var SuccessfulUploadCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "successful_upload_counter",
		Help: "count attempts to uploads to s3 bucket",
	},
)

// TxnLatencyGauge  transaction latency/response time
var TxnLatencyGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "latency_gauge",
		Help: "gauge time it takes to send a response",
	},
)

// DataIngestionGauge  real-time data ingestion delay/lag
var DataIngestionGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "data_ingestion_gauge",
		Help: "gauge time it takes to receive a response",
	},
)

var Endpointcounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "endpoint_counter_vec",
		Help: "counts all endpoint hits",
	},
	[]string{"path"},
)

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}
