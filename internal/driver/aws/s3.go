package aws

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var (
	AWS_S3_BUCKET_NAME = os.Getenv("AWS_S3_BUCKET_NAME")
	AWS_REGION         = os.Getenv("AWS_REGION")
)

type Bucket struct {
	S3Client *s3.Client
}

func (b *Bucket) Upload(file multipart.File, key string) error {

	_, err := b.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET_NAME),
		Key:    aws.String(key),
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		log.Printf("Couldn't upload file. Here's why: %v\n", err)
		return err
	}

	return nil
}
