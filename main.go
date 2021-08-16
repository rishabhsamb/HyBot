package main

import (
	"log"

	"os"

	"os/signal"

	"syscall"

	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("Hello")
	token := "ODc2NDgyMzUzMTM5MTY3MjMy.YRktzQ.DWvYcEteGk6MdXKMvlnEUHsLTM8"
	b, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err.Error())
	}

	err = b.Open()
	if err != nil {
		log.Panic("Could not connect to Discord", err)
		return
	}

	log.Print("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	b.Close()

}
