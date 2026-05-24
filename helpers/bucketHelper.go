package helpers

import (
	"s3-file-manager/config"
	"s3-file-manager/models"
)

func mapBucketData(requestbody models.S3Request) config.BucketInfo {
	for _, bucket := range config.Config.BucketInfo {
		if bucket.Identifier == requestbody.Identifier {
			switch requestbody.Identifier {
			case "HL7":
				bucket.Directory = "" //build directory in required format

			}
			return bucket
		}

	}
	return config.BucketInfo{}
}
