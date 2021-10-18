package main

import (
	"context"
	"log"

	"os/signal"

	"os"

	"syscall"

	"github.com/joho/godotenv"

	"github.com/bwmarrin/discordgo"
)

var ob = OutburstHandlerStruct{loadOutbursts("outbursts.txt"), "outbursts.txt"}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	firestoreClient := createClient(ctx)

	token := os.Getenv("TOKEN")
	b, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err.Error())
	}
	b.AddHandler(commandHandler)

	err = b.Open()
	if err != nil {
		log.Panic("Could not connect to Discord", err)
		return
	}

	defer b.Close()
	defer ob.saveOutbursts()

	log.Print("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
