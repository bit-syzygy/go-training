package slackBotServer

import (
	"github.com/nlopes/slack"
)

//Handler is an interface that defines how to handle a slack message event
type Handler interface {
	Handle(msg *slack.MessageEvent) string
}

//Matcher is an interface that tests a slack message
type Matcher interface {
	Match(*slack.MessageEvent) bool
}

//HandleChannel defines a chanel that messages can be fed into in order for replys to be generated
func HandleChannel(matcher Matcher, handler Handler, reply chan<- Reply) chan<- *slack.MessageEvent {
	c := make(chan *slack.MessageEvent)

	go func() {
		for {
			msg := <-c
			go handle(msg, matcher, handler, reply)
		}
	}()

	return c
}

func handle(msg *slack.MessageEvent, matcher Matcher, handler Handler, replyChannel chan<- Reply) {
	if !matcher.Match(msg) {
		return
	}

	reply := Reply{
		MessageText: handler.Handle(msg),
		ReplyTo:     msg,
	}

	replyChannel <- reply
}
