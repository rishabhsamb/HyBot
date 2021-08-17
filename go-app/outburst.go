package main

import (
	"bufio"

	"log"

	"os"

	"strings"

	"strconv"

	"github.com/bwmarrin/discordgo"
)

type outburst interface {
	fire(s *discordgo.Session, cid string)
	getKey() string
}

type OutburstHandlerStruct struct {
	outburstSlice []outburst
}

func (ob *OutburstHandlerStruct) OutburstHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, v := range ob.outburstSlice {
		if v.getKey() == m.Content {
			v.fire(s, m.ChannelID)
		}
	}
}

type vanillaOutburst struct {
	key      string
	count    uint64
	messages []string
}

func (v *vanillaOutburst) getKey() string {
	return v.key
}

func (v *vanillaOutburst) fire(s *discordgo.Session, cid string) {
	for _, str := range v.messages {
		s.ChannelMessageSend(cid, str)
	}
	s.ChannelMessageSend(cid, v.key+" has been called "+strconv.FormatUint(v.count, 10)+" times")
}

func loadOutbursts(filepath string) []outburst {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("FAILED TO READ FILE WITH PATH " + filepath)
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // WE ASSUME EACH LINE IS LESS THAN 65536 CHARACTERS

	var retSlice []outburst
	for scanner.Scan() {
		switch lineSlice := strings.Split(scanner.Text(), " || "); {
		case lineSlice[0] == "vo":
			count, err := strconv.ParseUint(lineSlice[2], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			retSlice = append(retSlice, &vanillaOutburst{lineSlice[1], count, lineSlice[3:]})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return retSlice
}
