package models

//import v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"

type S3Response struct {
	StatusCode   int    `json:"statusCode"`
	Message      string `json:"message"`
	FileContents []byte `json:"fileContents,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
	FileName     string `json:"fileName,omitempty"`
	//PresignedUrl v4.PresignedHTTPRequest `json:"presignedUrl,omitempty"`
}

type S3Request struct {
	H17type      string `json:"h17Type,omitempty"`
	FileContents []byte `json:"fileContents,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
	FileName     string `json:"fileName,omitempty"`
}
