package s3

import (
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (h *S3Handler) Clean() {
	sess := getSession(h.S3)
	s3Svc := s3.New(sess)
	go func(s3Svc *s3.S3, bucket string, logger *slog.Logger) {
		defer wg.Done()
		deleteObject(s3Svc, h.S3.Bucket, h.Logger)
	}(s3Svc, h.S3.Bucket, h.Logger)

	wg.Wait()
}

func deleteObject(s3Svc *s3.S3, bucket string, logger *slog.Logger) {
	objects, err := s3Svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		logger.Error("Failed to list objects", slog.String("error", err.Error()))
	}
	for _, object := range objects.Contents {
		s3Svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    object.Key,
		})
	}
}
