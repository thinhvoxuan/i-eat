package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	godotenv.Load()
	bot := slacker.NewClient(os.Getenv("TOKEN"))

	bot.Command("Menu", "List of good today", func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("List of good today")
	})

	bot.Command("Add <number>", "Add food", func(request slacker.Request, response slacker.ResponseWriter) {
		number := request.IntegerParam("number", 1)
		response.Reply(fmt.Sprintf("Add %d", number))
	})

	bot.DefaultCommand(func(request slacker.Request, response slacker.ResponseWriter) {
		// bot.defaultHelp
		response.Reply("type `help @Cereals`")
	})

	bot.Command("Wifi", "Password", func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("123444222")
	})

	bot.Command("Remove <number>", "Remove food", func(request slacker.Request, response slacker.ResponseWriter) {
		number := request.IntegerParam("number", 1)
		response.Reply(fmt.Sprintf("Remove %d", number))
	})

	bot.Command("Reset", "Remove all foods", func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("Remove all foods")
	})

	bot.Command("Review", "Review your food", func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("Review your food")
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
