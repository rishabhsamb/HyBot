package main

import (
	"log"

	"os"

	"os/signal"

	"syscall"

	"github.com/bwmarrin/discordgo"
)

var ob = OutburstHandlerStruct{loadOutbursts("outbursts.txt")}

func main() {
	token := "ODc2NDgyMzUzMTM5MTY3MjMy.YRktzQ.DWvYcEteGk6MdXKMvlnEUHsLTM8"
	b, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err.Error())
	}
	b.AddHandler(ob.OutburstHandler)

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
