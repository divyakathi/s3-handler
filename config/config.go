package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Configuration struct {
	Debug         bool         `json:"debug"`
	ConfigVersion string       `json:"configVersion"`
	Environment   string       `json:"env"`
	DbUtilUrl     string       `json:"dbUtilUrl"`
	ProducerUrl   string       `json:"producerUrl"`
	BucketInfo    []BucketInfo `json:"bucketInfo"`
}

type BucketInfo struct {
	Region       string `json:"region"`
	BucketName   string `json:"bucketName"`
	Directory    string `json:"directory"`
	AccessKey    string `json:"accessKey"`
	SecretKey    string `json:"secretKey"`
	SessionToken string `json:"sessionToken"`
	Identifier   string `json:"identifier"`
}

var Config Configuration
var err error
var TestConfig = Configuration{
	DbUtilUrl:   "localhost:0000",
	ProducerUrl: "localhost:1111",
	BucketInfo: []BucketInfo{
		{
			Region:     "us-west-1",
			BucketName: "test-bucket",
			AccessKey:  "123",
			SecretKey:  "2344",
			Identifier: "SES",
		},
	},
}

func init() {
	Config, err = GetConfig()
	if err != nil {
		log.Fatal("error getting app config: ", err)
	}
}

func GetConfig() (Configuration, error) {
	if strings.HasSuffix(os.Args[0], ".test") {
		return TestConfig, nil
	}
	absPath, _ := filepath.Abs("config/config.json")
	config := Configuration{}
	properties, err := os.Open(absPath)
	if err != nil {
		return config, err
	}
	decoder := json.NewDecoder(properties)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return config, err
	}
	return configuration, nil
}
