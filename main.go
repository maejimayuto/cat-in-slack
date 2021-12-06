package main

import (
	"cat-in-slack/cat"
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	result := cat.GetCat()
	fmt.Printf("%s", result)
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	if _, _, err := api.PostMessage("04-cat", slack.MsgOptionText("test", false)); err != nil {
		log.Println(err)
		return
	}
}
