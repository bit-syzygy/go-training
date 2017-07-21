package slackBotServer

import (
	"testing"

	"github.com/nlopes/slack"
)

func TestReplyChannel(t *testing.T) {
	api := TestAPI{}
	c := ReplyChannel(api)
	m := slack.MessageEvent{}
	m.Channel = "TestChannel"
	reply := Reply{
		MessageText: "Some text",
		ReplyTo:     &m,
	}
	c <- reply
}
