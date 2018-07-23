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


func main() {
	var name = os.Args[1]
	var args = os.Args[2:]

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

	err = cmd.Wait()

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