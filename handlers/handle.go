package handlers

import (
	"encoding/json"
	"net/http"
	"s3-file-manager/helpers"
	"s3-file-manager/models"
	"strconv"
)

var reStartApp = 0

func HealthCheck(w http.ResponseWriter, request *http.Request) {
	if reStartApp == 0 {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("s3-file-handler is running"))
		return
	}
}

func RestartApp(w http.ResponseWriter, request *http.Request) {
	reStartApp = 1
	var msg = "Successfully updated Health Check response to " + strconv.Itoa(reStartApp)
	w.WriteHeader(200)
	_, _ = w.Write([]byte(msg))
	return

}

func ProcessFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bucketResponse := models.S3Response{}
		switch r.Method {
		case http.MethodPost:

			bucketResponse = helpers.PostFileToS3(r)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(bucketResponse.StatusCode)
		_ = json.NewEncoder(w).Encode(bucketResponse)
	}
}
