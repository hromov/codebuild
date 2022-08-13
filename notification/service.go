package notification

import "log"

type TopicSender interface {
	Send(msg string) error
}

type service struct {
	topicSender TopicSender
}

func New(ts TopicSender) *service {
	return &service{ts}
}

func (s *service) Notify(msg string) {
	if err := s.topicSender.Send(msg); err != nil {
		log.Println("Error while sending msg to topic: ", err.Error())
	}
}
