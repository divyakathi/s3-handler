package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	log "s3-file-manager/logging"
	"s3-file-manager/models"
	"s3-file-manager/services"
)

func PostFileToS3(request *http.Request) models.S3Response {
	//get request body
	requestBody := getRequestBody(request)
	logFields := log.LogFields{
		FileName:   requestBody.FileName,
		Identifier: requestBody.Identifier,
	}
	log.Info(logFields, "request to save file received")
	//send to bucket
	bucketData := mapBucketData(requestBody)
	return services.SendToS3(bucketData, requestBody.FileContents, logFields)

}

func getRequestBody(r *http.Request) models.S3Request {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(log.LogFields{}, "error reading body: ", err)
		return models.S3Request{}
	}
	s3Request := models.S3Request{}
	if err := json.Unmarshal(body, &s3Request); err != nil {
		return models.S3Request{}
	}
	return s3Request
}
