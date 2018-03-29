package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("APITELEGRAM"))
	bot.Debug = false

	if err != nil {
		log.Panic(err)
	}

	log.Println(bot.Self.UserName)

	consolescanner := bufio.NewScanner(os.Stdin)
	if err := consolescanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for consolescanner.Scan() {
		input := consolescanner.Text()
		fmt.Println("Sending to telegram: ", input)
		msg := tgbotapi.NewMessage(os.Getenv("CHATIDTELEGRAM"), input)
		bot.Send(msg)
	}

}
