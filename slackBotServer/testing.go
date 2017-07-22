package slackBotServer

import (
	"github.com/nlopes/slack"
)

//TestAPI is a mocked slack api for testing
type TestAPI struct {
	PostChannel chan TestPost
}

//TestPost is contains the infomation in a slack post
type TestPost struct {
	ChannelID   string
	MessageText string
	Params      slack.PostMessageParameters
}

//PostMessage mocks the slack api post message
func (t TestAPI) PostMessage(channelID string, messageText string, params slack.PostMessageParameters) (string, string, error) {
	t.PostChannel <- TestPost{
		ChannelID:   channelID,
		MessageText: messageText,
		Params:      params,
	}
	return "", "", nil
}

//NewRTM mocks the slack rtm maker
func (t TestAPI) NewRTM() *slack.RTM {
	return &slack.RTM{}
}

//TestMatch impliments the Matcher interface
type TestMatch struct {
}

//Match allows TestMatch to impliment the Matcher interface
func (m TestMatch) Match(*slack.MessageEvent) bool {
	return true
}

//TestHandle impliments the Handler interface
type TestHandle struct{}

//Handle allows TestHandler to impliment the Handler interface
func (h TestHandle) Handle(*slack.MessageEvent) string {
	return "A response"
}
