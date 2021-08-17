package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type command interface {
	getPrefix() string
	fire(args []string, s *discordgo.Session, cid string)
}

type addOutburst struct {
	prefix string
}

func (addOb *addOutburst) getPrefix() string {
	return addOb.prefix
}

func (addOb *addOutburst) fire(args []string, s *discordgo.Session, cid string) {
	if len(args) < 3 {
		s.ChannelMessageSend(cid, "correct usage is ¿ addob key=\"<key>\" \"<message1>\" \"<message2>\" ... \"<messagen>\"")
		return
	}

	var key string
	if strings.HasPrefix(args[1], "key=\"") && strings.HasSuffix(args[1], "\"") {
		key = args[1][5 : len(args[1])-1]
		//s.ChannelMessageSend(cid, key)
	} else {
		s.ChannelMessageSend(cid, "key is poorly formatted")
		return
	}

	var messageSlice []string
	for i := 2; i < len(args); i++ {
		if strings.HasPrefix(args[i], "\"") && strings.HasSuffix(args[i], "\"") {
			messageSlice = append(messageSlice, args[i][1:len(args[i])-1])
			fmt.Println(messageSlice)
		} else {
			s.ChannelMessageSend(cid, "message "+strconv.FormatInt(int64(i), 10)+" is poorly formatted")
			return
		}
	}
	ob.addOutburst(&vanillaOutburst{key, 0, messageSlice})
}

var commands = []command{&addOutburst{"addob"}}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "?") {
		ob.OutburstHandler(s, m)
		return
	}

	args := strings.Split(m.Content, " ")[1:]
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, "?")
		return
	}

	for i := 0; i < len(commands); i++ {
		if args[0] == commands[i].getPrefix() {
			commands[i].fire(args, s, m.ChannelID)
			return
		}
	}
}
