package services

import (
	"bytes"
	"context"
	"s3-file-manager/config"
	log "s3-file-manager/logging"
	"s3-file-manager/models"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//var s3Client *s3.Client

type Uploader interface {
	Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error)
}

func SendToS3(bucketInfo config.BucketInfo, fileContent []byte, logFields log.LogFields) models.S3Response {
	bucketResponse := models.S3Response{}
	log.Info(logFields, "REGION: ", bucketInfo.Region, " IDENTIFIER: ", bucketInfo.Identifier)
	creds := credentials.NewStaticCredentialsProvider(bucketInfo.AccessKey, bucketInfo.SecretKey, bucketInfo.SessionToken)
	cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithCredentialsProvider(creds), awsConfig.WithRegion(bucketInfo.Region))

	if err != nil {
		msg := "error creating aws config: " + err.Error()
		bucketResponse = mapBucketResponse(500, msg, bucketInfo.Identifier, logFields.FileName, fileContent)
		return bucketResponse
	}
	s3Client := s3.NewFromConfig(cfg)
	bucket := bucketInfo.BucketName
	fileName := bucketInfo.Directory + logFields.FileName
	r := bytes.NewReader(fileContent)

	uploader := manager.NewUploader(s3Client)
	input := s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName), //// S3 object key (path in bucket)
		Body:   r,
	}
	_, err = upload(uploader, &input)
	bucketResponse.StatusCode = 200
	bucketResponse.Message = "successfully uploaded to s3 bucket: " + bucketInfo.BucketName
	return bucketResponse

}

func upload(uploader Uploader, input *s3.PutObjectInput) (*manager.UploadOutput, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	output, err := uploader.Upload(ctx, input)
	return output, err
}

func mapBucketResponse(statusCode int, msg string, identifier string, fileName string, file []byte) models.S3Response {
	s3Resp := models.S3Response{
		StatusCode:   statusCode,
		Message:      msg,
		FileName:     fileName,
		FileContents: file,
		Identifier:   identifier,
	}
	return s3Resp
}
