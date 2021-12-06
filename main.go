package main

import (
	"cat-in-slack/cat"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		return
	}
	result := cat.GetCat()
	fmt.Printf("%s", result)
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	if _, _, err := api.PostMessage("04-cat", slack.MsgOptionText("test", false)); err != nil {
		log.Println(err)
		return
	}
}
