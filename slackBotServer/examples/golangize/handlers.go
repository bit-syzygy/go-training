package main

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

//GoMatch impliments slackBotServer.Matcher
type GoMatch struct{}

//Match tests if the message had "go" in it
func (g GoMatch) Match(msg *slack.MessageEvent) bool {
	if strings.Contains(msg.Text, "go") || strings.Contains(msg.Text, "Go") || strings.Contains(msg.Text, "GO") {
		return true
	}
	return false
}

//GoHandle impliments slackBotServer.Handler
type GoHandle struct{}

//Handle replaces go with golang
func (h GoHandle) Handle(msg *slack.MessageEvent) string {
	text := msg.Text
	text = strings.Replace(text, "go", "golang", -1)
	text = strings.Replace(text, "Go", "Golang", -1)
	text = strings.Replace(text, "GO", "GOLANG", -1)
	reply := fmt.Sprintf("Did you mean \"%s\" ?", text)
	fmt.Printf("received \"%s\" sending reply \"%s\" \n", msg.Text, reply)
	return reply
}
