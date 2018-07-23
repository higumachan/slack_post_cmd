package main

import (
	"os"
	"os/exec"

	"github.com/nlopes/slack"
	"fmt"
	"log"
	"bufio"
	"io"
)


func PostSlack(programName string, commandError error)  {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	username := os.Getenv("SLACK_USER_ID")
	var message string
	if commandError == nil {
		message = fmt.Sprintf("<@%s> finished %s", username, programName)
	} else {
		message = fmt.Sprintf("<@%s> failed %s %s", username, programName, commandError)
	}
	api.PostMessage(os.Getenv("SLACK_CHANNEL_ID"), message, slack.PostMessageParameters{})
}

func RunCommandAndStreamOutputStdout(name string, args []string) error {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	rd := bufio.NewReader(stdout)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	defer stdout.Close()
	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(line))
	}

	return cmd.Wait()
}

func main() {
	var name = os.Args[1]
	var args = os.Args[2:]

	err := RunCommandAndStreamOutputStdout(name, args)
	PostSlack(name, err)
}