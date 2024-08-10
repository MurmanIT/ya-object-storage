package s3

import (
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (h *S3Handler) Echo() []string {
	sess := getSession(h.S3)
	s3Svc := s3.New(sess)
	listObject := []string{}

	objects, err := s3Svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(h.S3.Bucket),
	})

	if err != nil {
		h.Logger.Error("Failed to list objects", slog.String("error", err.Error()))
	}

	for _, object := range objects.Contents {
		listObject = append(listObject, *object.Key)
	}
	return listObject
}
