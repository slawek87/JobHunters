package utils

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"bytes"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/slawek87/JobHunters/conf"
	"os"
)


func UploadToS3(file *os.File) error {
	config := conf.S3Config()
	svc := s3.New(session.New(), config)

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/media/" + file.Name()
	params := &s3.PutObjectInput{
		Bucket: aws.String("testBucket"),
		Key: aws.String(path),
		Body: fileBytes,
		ContentLength: aws.Int64(size),
		ContentType: aws.String(fileType),
	}
	_, err := svc.PutObject(params)
	if err != nil {
		return err
	}
}