package conf

import (
	"gopkg.in/mgo.v2"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
)

// Returns MongoDB session with selected DB.
func MongoDB() (*mgo.Session, *mgo.Database) {
	session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	return session, session.DB("JobHunters")
}


const AWS_ACCESS_KEY_ID = "AKIAJ2765YRNN5XTXYGQ"
const AWS_SECRET_ACCESS_KEY = "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"
const AWS_REGION = "eu-west-1"
const S3_BUCKET_NAME = "jobhunters"


// Returns AWS S3 config.
func S3Config() *aws.Config {
	creds := credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, "")

	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}

	return aws.NewConfig().WithRegion(AWS_REGION).WithCredentials(creds)
}