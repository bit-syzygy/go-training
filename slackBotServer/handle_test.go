package slackBotServer

import (
	"fmt"
	"testing"

	"github.com/nlopes/slack"
)

func TestHandleChannel(t *testing.T) {
	msg := &slack.MessageEvent{}
	r := make(chan Reply)

	go func() {
		for {
			reply := <-r
			fmt.Println(reply)
		}
	}()

	h := HandleChannel(TestMatch{}, TestHandle{}, r)

	h <- msg
}
