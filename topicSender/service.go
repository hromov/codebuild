package topicsender

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type service struct {
	svc      *sns.SNS
	topicARN string
}

func New(topicARN string) *service {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	return &service{svc, topicARN}
}

func (s *service) Send(msg string) error {
	input := sns.PublishInput{
		Message:  aws.String(msg),
		TopicArn: aws.String(s.topicARN),
	}
	_, err := s.svc.Publish(&input)
	return err
}
