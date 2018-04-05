package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("APITELEGRAM"))
	bot.Debug = false

	chatid := flag.Int64("chatid", 0, "your telegram chatid")

	flag.Parse()

	if err != nil {
		log.Panic(err)
	}

	if len(os.Args) > 2 {
		log.Println(bot.Self.UserName)
		input := strings.Join(os.Args[2:], " ")
		msg := tgbotapi.NewMessage(*chatid, input)
		fmt.Println("Sending to telegram: ", input)
		bot.Send(msg)
	} else {

		consolescanner := bufio.NewScanner(os.Stdin)
		if err := consolescanner.Err(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		for consolescanner.Scan() {
			log.Println(bot.Self.UserName)
			input := consolescanner.Text()
			msg := tgbotapi.NewMessage(*chatid, input)
			fmt.Println("Sending to telegram: ", input)
			bot.Send(msg)
		}
	}
}
