package s3

import (
	"ya-storage/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getSession(ConfigS3 *config.S3) *session.Session {
	defaultResolver := endpoints.DefaultResolver()
	s3ResolverFn := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == s3.ServiceID && region == ConfigS3.Region {
			return endpoints.ResolvedEndpoint{
				PartitionID:   "yc",
				URL:           ConfigS3.Url,
				SigningRegion: ConfigS3.Region,
			}, nil
		}
		return defaultResolver.EndpointFor(service, region, optFns...)
	}
	session := session.Must(session.NewSession(&aws.Config{
		Endpoint:    aws.String(ConfigS3.Url),
		Region:      aws.String(ConfigS3.Region),
		Credentials: credentials.NewStaticCredentials(ConfigS3.AccessKey, ConfigS3.SecretKey, ""),
		EndpointResolver: endpoints.ResolverFunc(
			s3ResolverFn,
		),
	}))
	return session
}
