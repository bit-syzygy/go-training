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
	if strings.Contains(msg.Text, "go") {
		return true
	}
	return false
}

//GoHandle impliments slackBotServer.Handler
type GoHandle struct{}

//Handle replaces go with golang
func (h GoHandle) Handle(msg *slack.MessageEvent) string {
	reply := fmt.Sprintf("Did you mean \"%s\" ?", strings.Replace(msg.Text, "go", "golang", -1))
	fmt.Printf("received \"%s\" sending reply \"%s\" \n", msg.Text, reply)
	return reply
}
