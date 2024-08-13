package s3

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (h *S3Handler) UploadFile(pathFile string, key string) string {
	file, err := os.Open(pathFile)
	if err != nil {
		h.Logger.Error("Failed to open file", slog.String("error", err.Error()))
	}
	defer file.Close()

	wg.Add(1)

	ch := make(chan string)

	go func() {
		defer wg.Done()
		h.uploadFiles(key, file, ch)
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	link := <-ch
	return link
}
func (h *S3Handler) uploadFiles(key string, file *os.File, ch chan string) {
	sess := getSession(h.S3)
	s3Svc := s3.New(sess)
	_, err := s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(h.S3.Bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		fmt.Println("Failed to upload file", err)
	}
	req, _ := s3Svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(h.S3.Bucket),
		Key:    aws.String(key),
	})
	u, _, err := req.PresignRequest(60 * time.Minute)
	if err != nil {
		h.Logger.Error("Failed to generate link", slog.String("error", err.Error()))
	}
	ch <- u
}
