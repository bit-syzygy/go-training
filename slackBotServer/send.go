package slackBotServer

import (
	"fmt"

	"github.com/nlopes/slack"
)

//Reply is sruct that describes a slack reply
type Reply struct {
	MessageText string
	ReplyTo     *slack.MessageEvent
}

//PostMessager defines an interface for posting to slack
type PostMessager interface {
	PostMessage(channelID string, messsageText string, params slack.PostMessageParameters) (string, string, error)
}

//ReplyChannel creates a reply chanel for a bot with the given oauth token
//a Reply pushed to this channel will be sent
func ReplyChannel(api PostMessager) chan<- Reply {
	c := make(chan Reply)

	go func() {
		for {
			reply := <-c
			params := slack.PostMessageParameters{ThreadTimestamp: reply.ReplyTo.EventTimestamp}
			_, _, err := api.PostMessage(reply.ReplyTo.Channel, reply.MessageText, params)
			if err != nil {
				fmt.Println("Warning: Unable to post message")
			}
		}
	}()

	return c
}
