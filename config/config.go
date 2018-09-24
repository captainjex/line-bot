package config

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

func ConnectLineBot() *linebot.Client {
	client, err := linebot.New(
		"dafca8e4125c3547f7401d6091f8f783",
		"3nQXlYtVHy6GMlVX+MNhgpkTOixTXtJCsEUezJMOr5jyur5RamifYwiy9mFVDJIpNW8ii4BVQjtT6sRcGda7vA0AO5kjMe6SQDpot1mqy/uc4u3jzX6XTwbczXX+cfCz1cLeph5h/dG/qQp014ItUAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
