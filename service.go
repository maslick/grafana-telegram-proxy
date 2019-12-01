package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ITelegramSender interface {
	SendTelegramMessage(message string) ([]byte, error)
}

type TelegramSender struct{}

func (t *TelegramSender) SendTelegramMessage(message string) ([]byte, error) {
	token := getEnv("BOT_TOKEN", "")
	chatId := getEnv("CHAT_ID", "")

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?parse_mode=html", token)
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(Message{chatId, message})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json; charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	text, _ := ioutil.ReadAll(resp.Body)
	return text, err
}
