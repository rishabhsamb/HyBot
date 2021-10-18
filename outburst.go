package main

import (
	"bufio"

	"log"

	"os"

	"strings"

	"strconv"

	"math/rand"

	"github.com/bwmarrin/discordgo"

	"time"
)

type outburst interface {
	fire(s *discordgo.Session, cid string)
	getKey() string
	saveStringSlice() []string
}

type routburst interface {
	outburst
	getImageId() int
}

type vanillaOutburst struct {
	key      string
	count    uint64
	messages []string
}
type randomOutburst struct {
	base     vanillaOutburst
	picCount int
	images   []string
}

func (ro *randomOutburst) getImageId() int {
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	return r2.Intn(len(ro.images))
}

type OutburstHandlerStruct struct {
	outburstSlice []outburst
	saveFile      string
}

func (ob *OutburstHandlerStruct) OutburstHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, v := range ob.outburstSlice {
		if v.getKey() == m.Content {
			v.fire(s, m.ChannelID)
		}
	}
}

func (ob *OutburstHandlerStruct) addOutburst(toAdd outburst) {
	ob.outburstSlice = append(ob.outburstSlice, toAdd)
}

func (ob *OutburstHandlerStruct) saveOutbursts() {
	f, err := os.OpenFile(ob.saveFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for i := 0; i < len(ob.outburstSlice); i++ {
		_, err = f.WriteString(strings.Join(ob.outburstSlice[i].saveStringSlice(), " || ") + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (v *vanillaOutburst) getKey() string {
	return v.key
}
func (r *randomOutburst) getKey() string {
	return r.base.key
}

func (v *vanillaOutburst) fire(s *discordgo.Session, cid string) {
	for _, str := range v.messages {
		s.ChannelMessageSend(cid, str)
	}
	v.count++
	s.ChannelMessageSend(cid, v.key+" has been called "+strconv.FormatUint(v.count, 10)+" times")
}

func (r *randomOutburst) fire(s *discordgo.Session, cid string) {
	for _, str := range r.base.messages {
		s.ChannelMessageSend(cid, str)
	}
	s.ChannelMessageSend(cid, r.images[r.getImageId()])
	r.base.count++
	s.ChannelMessageSend(cid, r.base.key+" has been called "+strconv.FormatUint(r.base.count, 10)+" times")
}

func (v *vanillaOutburst) saveStringSlice() []string {
	return append([]string{"vo", v.key, strconv.FormatUint(v.count, 10)}, v.messages...)
}

func (r *randomOutburst) saveStringSlice() []string { //fix
	return append([]string{"ro", r.base.key, strconv.FormatUint(r.base.count, 10)}, r.base.messages...)
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
		case lineSlice[0] == "ro":
			pics := strings.Split(lineSlice[len(lineSlice)-1], " && ")
			count, err := strconv.ParseUint(lineSlice[2], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			retSlice = append(retSlice, &randomOutburst{vanillaOutburst{lineSlice[1], count, lineSlice[3 : len(lineSlice)-1]}, len(pics), pics}) //2nd last element
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return retSlice
}
