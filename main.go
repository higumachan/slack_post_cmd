package main

import (
	"os"
	"os/exec"

	"github.com/nlopes/slack"
	"fmt"
)


func main() {
	var name = os.Args[1]
	var args = os.Args[2:]

	cmd := exec.Command(name, args...)

	err := cmd.Run()
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	username := os.Getenv("SLACK_USER_ID")
	var message string
	if err == nil {
		message = fmt.Sprintf("<@%s> finished %s", username, name)
	} else {
		message = fmt.Sprintf("<@%s> failed %s %s", username, name, err)
	}
	api.PostMessage(os.Getenv("SLACK_CHANNEL_ID"), message, slack.PostMessageParameters{})
}