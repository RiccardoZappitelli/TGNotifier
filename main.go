package main

import (
	"encoding/json"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserInfo struct {
	Token  string `json:"token"`
	ChatId int64  `json:"chatid"`
}

func getUserInfo(filename string) (*UserInfo, error) {
	file, err := os.Open("config.json")
	if err != nil {
		//handle nigger
	}
	defer file.Close()

	var userinfo UserInfo
	decoder := json.NewDecoder(file)
	decoder.Decode(&userinfo)

	return &userinfo, nil
}

func main() {
	args := os.Args[1:]
	userinfo, err := getUserInfo("config.json")
	if err != nil {
		//handle error
	}

	token := userinfo.Token
	chatid := int64(userinfo.ChatId)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	msg := tgbotapi.NewMessage(chatid, strings.Join(args, " "))
	bot.Send(msg)
}
