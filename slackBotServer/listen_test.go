package slackBotServer

import (
	"testing"
	"time"

	"github.com/nlopes/slack"
)

func TestListen(t *testing.T) {
	api := TestAPI{}
	c := make(chan *slack.MessageEvent)
	go Listen(api, c)

	<-time.After(time.Millisecond * 100)
	t.SkipNow()
}
