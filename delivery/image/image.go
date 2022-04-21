package image

import (
	"capstone/configs"
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadImage(path string, fileName string, fileData multipart.File) (string, error) {
	// use the S3 uploader
	sess := configs.GetSession()

	// create an uploader
	uploader := s3manager.NewUploader(sess)

	// upload file to S3
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(path + "/" + fileName + ".jpeg"),
		Body:   fileData,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	return result.Location, nil
}

func CheckImageFileExtension(fileName string) (string, error) {
	extension := strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:])
	jpg := extension != "jpg"
	jpeg := extension != "jpeg"
	png := extension != "png"
	if jpg && jpeg && png {
		return "", fmt.Errorf("unsupported file type")
	}
	return extension, nil
}

func CheckImageFileSize(size int64) error {
	if size == 0 {
		return fmt.Errorf("file undetectable")
	}
	if size > 3000000 {
		return fmt.Errorf("file is too big")
	}
	return nil
}
