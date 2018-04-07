//TODO
// Better way to handle errors

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("APITELEGRAM"))
	bot.Debug = false

	if err != nil {
		log.Panic(err)
	}

	chatid, err := strconv.ParseInt(os.Getenv("CHATID"), 10, 64)

	if err != nil {
		log.Panic(err)
	}

	if len(os.Args) > 1 {
		input := strings.Join(os.Args[1:], " ")
		msg := tgbotapi.NewMessage(chatid, input)
		fmt.Println(bot.Self.UserName)
		fmt.Println("Sending to telegram: ", input)
		bot.Send(msg)
	} else {

		consolescanner := bufio.NewScanner(os.Stdin)
		if err := consolescanner.Err(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		for consolescanner.Scan() {
			input := consolescanner.Text()
			msg := tgbotapi.NewMessage(chatid, input)
			fmt.Println(bot.Self.UserName)
			fmt.Println("Sending to telegram: ", input)
			bot.Send(msg)
		}
	}
}
