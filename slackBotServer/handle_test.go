package slackBotServer

import (
	"testing"

	"github.com/nlopes/slack"
)

func TestHandleChannel(t *testing.T) {
	msg := &slack.MessageEvent{}
	r := make(chan Reply)

	go func() {
		for {
			reply := <-r
			text := TestHandle{}.Handle(nil)
			if reply.ReplyTo != msg || reply.MessageText != text {
				t.Error("TestHandleChannel: Unexpected Reply")
			}
		}
	}()

	h := HandleChannel(TestMatch{}, TestHandle{}, r)

	h <- msg
}
