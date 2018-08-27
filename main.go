package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	godotenv.Load()
	bot := slacker.NewClient(os.Getenv("TOKEN"))

	db, err := gorm.Open("postgres", "host=db user=postgres dbname=cereals password=example sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Println("Error-Connect-DB")
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "New Product", Price: 1000})

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

	log.Println("Complete loading")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errListen := bot.Listen(ctx)
	if errListen != nil {
		log.Fatal(err)
	}

}
