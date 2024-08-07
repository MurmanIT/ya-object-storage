package s3

import (
	"fmt"
	"log/slog"
	"os"
	"ya-storage/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (h *S3Handler) UploadFile(pathFile string, key string) {
	file, err := os.Open(pathFile)
	if err != nil {
		h.Logger.Error("Failed to open file", slog.String("error", err.Error()))
	}
	defer file.Close()

	wg.Add(1)

	go func() {
		defer wg.Done()
		uploadFiles(h.S3, key, file)
	}()
	wg.Wait()
}
func uploadFiles(ConfigS3 *config.S3, key string, file *os.File) {
	sess := getSession(ConfigS3)
	s3Svc := s3.New(sess)
	_, err := s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(ConfigS3.Bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		fmt.Println("Failed to upload file", err)
	}
}
