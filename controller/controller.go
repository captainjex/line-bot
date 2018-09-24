package controller

import (
	"net/http"
	"github.com/line/line-bot-sdk-go/linebot"
	"fmt"
	"github.com/ciazhar/line-bot/config"
	"log"
	"github.com/ciazhar/util"
)

func Base(w http.ResponseWriter, _ *http.Request)  {
	util.RespondWithJson(w,http.StatusOK, map[string]string{"message": "Welcome To Line Bot Service"})
}

func Callback(w http.ResponseWriter, req *http.Request)  {
	bot := config.ConnectLineBot()

	events, err := bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
				fmt.Println(message.Text)
			}
		}
	}
}