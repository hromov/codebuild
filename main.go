package main

import (
	"github.com/hromov/notification/notification"
	topicsender "github.com/hromov/notification/topicSender"
)

const topicARN = "fake_topic_arn"

type Notifier interface {
	Notify(msg string)
}

func main() {
	var ns Notifier
	ts := topicsender.New(topicARN)
	ns = notification.New(ts)
	ns.Notify("hello")
}
