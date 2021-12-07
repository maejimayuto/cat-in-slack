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
		log.Fatal(err)
	}
	cats := cat.GetCats()
	fmt.Printf("%s", cats[0].ToString())
	log.Printf("%s", cats[0].ToString())
	// TODO: 2MB 以上だと画像が展開されないので、別の画像に変更する #1
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	// TODO: 画像の block だけをいい感じに表示したい #2
	if _, _, err := api.PostMessage("04-cat", slack.MsgOptionText(cats[0].Url, false)); err != nil {
		log.Fatal(err)
	}
}
