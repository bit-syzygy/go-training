package main

import (
	"os"

	"github.com/bit-syzygy/go-training/slackBotServer"
	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	r := slackBotServer.ReplyChannel(api)
	h := slackBotServer.HandleChannel(GoMatch{}, GoHandle{}, r)
	slackBotServer.Listen(api, h)
}
