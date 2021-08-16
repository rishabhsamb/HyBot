package main

import (
	"log"

	"os"

	"os/signal"

	"syscall"

	"fmt"

	"github.com/bwmarrin/discordgo"

	"math/rand"

	"strconv"

	"time"
)

var hyoonCount = 0
var startTime string

func main() {
	fmt.Println("Hello")
	token := "ODc2NDgyMzUzMTM5MTY3MjMy.YRktzQ.DWvYcEteGk6MdXKMvlnEUHsLTM8"
	b, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err.Error())
	}

	b.AddHandler(HyoonHandler)

	err = b.Open()
	if err != nil {
		log.Panic("Could not connect to Discord", err)
		return
	} else {
		startTime = time.Now().Format("02-Jan-2006 15:04:05")
	}
	defer b.Close()

	log.Print("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func HyoonHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.ID == "724424874327539762" && rand.Intn(5) <= 2 {
		s.ChannelMessageSend(m.ChannelID, "Why don't you get some")
		s.ChannelMessageSend(m.ChannelID, "https://contentgrid.thdstatic.com/hdus/en_US/DTCCOMNEW/Articles/how-to-use-a-garden-hoe-step-1.jpg")
		return
	}

	if m.Content == "hyoon" {
		s.ChannelMessageSend(m.ChannelID, "https://pbs.twimg.com/media/Etp8tu7VgAAv4KG.jpg")
		var message = "Hyoon has been said " + strconv.Itoa(hyoonCount) + " time"
		if hyoonCount > 1 || hyoonCount == 0 {
			message = message + "s"
		}
		message = message + " since " + startTime
		s.ChannelMessageSend(m.ChannelID, message)
		hyoonCount = hyoonCount + 1
		return
	}

	if m.Content == "jimin" {
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/810539129279348746/872551080771342336/T2TtMax2RuCj47a4TWziebikpa8pd9rYQKvDehmZCXcPkZIbiAxPYmiZB8tiBE9_YiJxJHtFJaAAeQtWJxAfP6EvfwxFhyXJ4WGn.png")
		return
	}

	if m.Content == "drip" {
		s.ChannelMessageSend(m.ChannelID, "its tsunami szn :imp:")
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/875459399014043648/876665762343972884/deepfry.png")
		return
	}

	if m.Content == "rishabh" {
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/875459399014043648/876672643582013480/image0.png")
		return
	}
}
