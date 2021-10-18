package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	firestoreClient := createClient(ctx)
	var (
		commander commandHandler
	)

	commander.init(ctx, firestoreClient)

	token := os.Getenv("DISCORD_TOKEN")
	b, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err.Error())
	}
	b.AddHandler(commander.driver)

	err = b.Open()
	if err != nil {
		log.Panic("Could not connect to Discord", err)
		return
	}

	defer b.Close()

	log.Print("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
