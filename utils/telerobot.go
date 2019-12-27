package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func SendChat(message string) {
	bot, err := tgbotapi.NewBotAPI("811961426:AAHcO9SleJrSvcYbz7z5X2gTpbC_WOl1t4E")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	bot.GetUpdatesChan(u)

	msg := tgbotapi.NewMessage(1033112187, message)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}

}
