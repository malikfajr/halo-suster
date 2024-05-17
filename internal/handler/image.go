package handler

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/internal/driver/aws"
	"github.com/malikfajr/halo-suster/internal/exception"
)

var (
	AWS_S3_BUCKET_NAME = os.Getenv("AWS_S3_BUCKET_NAME")
)

type ImageHandler struct{}

func (i *ImageHandler) UploadImage(e echo.Context) error {
	file, err := e.FormFile("file")
	if err != nil {
		return e.JSON(400, exception.NewBadRequest("file required"))
	}

	// Check file size (10MB = 10 * 1024 * 1024 bytes)
	if file.Size > 2*1024*1024 || file.Size < 2*1024{
		return e.JSON(400, exception.NewBadRequest("File size min 2kb, max 2MB limit"))
	}

	ext := strings.Split(file.Filename, ".")
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	key := id.String() + "." + ext[len(ext)-1]

	src, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer src.Close()

	if valid, _ := i.isValidImage(src); valid == false {
		return e.JSON(400, exception.NewBadRequest(""))
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	bucket := &aws.Bucket{S3Client: client}
	bucket.Upload(src, key)

	if err != nil {
		log.Printf("Couldn't upload file. Here's why: %v\n", err)
		return e.JSON(500, "")
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", AWS_S3_BUCKET_NAME, key)

	return e.JSON(200, &jsonOk{
		Message: "success",
		Data: map[string]string{
			"imageUrl": url,
		},
	})

}

func (i *ImageHandler) isValidImage(file multipart.File) (bool, error) {
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return false, err
	}
	// Reset the read pointer of the file
	if _, err := file.Seek(0, 0); err != nil {
		return false, err
	}
	filetype := http.DetectContentType(buffer)
	return filetype == "image/jpeg" || filetype == "image/jpg", nil
}
