package main

import (
	"os"
	"os/exec"

	"fmt"

	"strings"

	"github.com/nlopes/slack"
)

func PostSlack(programName string, args []string, commandError error) {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	username := os.Getenv("SLACK_USER_ID")
	args_string := strings.Join(args, " ")
	var message string
	if commandError == nil {
		message = fmt.Sprintf("<@%s> finished %s %s", username, programName, args_string)
	} else {
		message = fmt.Sprintf("<@%s> failed %s %s %s", username, programName, args_string, commandError)
	}
	api.PostMessage(os.Getenv("SLACK_CHANNEL_ID"), message, slack.PostMessageParameters{})
}

func RunCommandAndStreamOutputStdout(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	var name = os.Args[1]
	var args = os.Args[2:]

	err := RunCommandAndStreamOutputStdout(name, args)
	PostSlack(name, args, err)
}
