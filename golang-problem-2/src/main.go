package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
	Url          string `json:"url"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	Keyboard       [][]KeyboardButton       `json:"keyboard"`
	ResizeKeyboard bool                     `json:"resize_keyboard"`
	OnTimeKeyboard bool                     `json:"one_time_keyboard"`
	Selective      bool                     `json:"selective"`
}

type SendMessage struct {
	ChatID      interface{} `json:"chat_id"`
	Text        string      `json:"text"`
	ParseMode   string      `json:"parse_mode"`
	ReplyMarkup interface{} `json:"reply_markup"`
}

func checkChatID(message *SendMessage) error {
	if message.ChatID == nil {
		return errors.New("chat_id is empty")
	}
	switch message.ChatID.(type) {
	case string:
		break
	case float64:
		message.ChatID = fmt.Sprint(message.ChatID)
	default:
		return errors.New("chat_id is unknown")
	}
	return nil
}

func ReadSendMessageRequest(fileName string) (*SendMessage, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var message SendMessage
	err = json.Unmarshal(file, &message)
	if err != nil {
		return nil, err
	}

	err = checkChatID(&message)
	if err != nil {
		return nil, err
	}

	if message.Text == "" {
		return nil, errors.New("text is empty")
	}

	var replyMarkupMap map[string]interface{}
	switch message.ReplyMarkup.(type) {
	case string:
		err = json.Unmarshal([]byte(message.ReplyMarkup.(string)), &replyMarkupMap)
	default:
		replyMarkupMap = message.ReplyMarkup.(map[string]interface{})
	}
	if err != nil {
		return nil, err
	}
	message.ReplyMarkup = replyMarkupMap

	var replyMarkup ReplyMarkup
	replyMarkupJson, _ := json.Marshal(message.ReplyMarkup)
	err = json.Unmarshal(replyMarkupJson, &replyMarkup)
	message.ReplyMarkup = replyMarkup

	return &message, nil
}

func main() {
	msg, err := ReadSendMessageRequest("input_sample2.json")
	//msg, err := ReadSendMessageRequest("input4.json")
	fmt.Println("######")
	fmt.Printf("%+v\n", msg)
	fmt.Println(err)
}
