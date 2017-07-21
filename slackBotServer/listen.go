package slackBotServer

import (
	"github.com/nlopes/slack"
)

//NewRTMer is an interface that impliments NewRTM in the way that the slack api does
type NewRTMer interface {
	NewRTM() *slack.RTM
}

//Listen runs a slack bot listening for message events
// token is the OAUTH token for the bot and messageChannel is the channel that messages are pushed to when received by the listner
func Listen(api NewRTMer, messageChannel chan<- *slack.MessageEvent) {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			info := rtm.GetInfo()
			if ev.User != "" && ev.User != info.User.ID {
				messageChannel <- ev
			}
		}
	}
}
