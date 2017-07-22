package slackBotServer

import (
	"testing"

	"github.com/nlopes/slack"
)

func TestReplyChannel(t *testing.T) {
	p := make(chan TestPost)
	api := TestAPI{PostChannel: p}
	c := ReplyChannel(api)
	m := slack.MessageEvent{}
	m.Channel = "TestChannel"
	reply := Reply{
		MessageText: "Some text",
		ReplyTo:     &m,
	}
	go func() {
		post := <-p
		if post.ChannelID != m.Channel || post.MessageText != reply.MessageText || post.Params.ThreadTimestamp != m.Timestamp {
			t.Error("TestReplyChannel: Incorrect message recieved")
		}
	}()
	c <- reply
}
