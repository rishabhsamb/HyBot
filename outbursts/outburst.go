package outbursts

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"

	"strconv"
)

type outburst struct {
	key            string   `firestore:"key"`
	callCount      uint64   `firestore:"callCount"`
	messages       []string `firestore:"messages"`
	randomMessages []string `firestore:"randomMessages"`
}

func (v *outburst) getKey() string {
	return v.key
}

func (v *outburst) getRandomMessageIndex() int {
	randomSeed := rand.NewSource(time.Now().UnixNano())
	randomEngine := rand.New(randomSeed)
	return randomEngine.Intn(len(v.randomMessages))
}

func (v *outburst) sendMessages(s *discordgo.Session, cid string) {
	for _, str := range v.messages {
		s.ChannelMessageSend(cid, str)
	}
	if len(v.randomMessages) > 0 {
		s.ChannelMessageSend(cid, v.randomMessages[v.getRandomMessageIndex()])
	}
}

func (v *outburst) updateCount(s *discordgo.Session, cid string) {
	v.callCount++
	s.ChannelMessageSend(cid, v.key+" has been called "+strconv.FormatUint(v.callCount, 10)+" times.")
}

func (v *outburst) fire(s *discordgo.Session, cid string) {
	v.sendMessages(s, cid)
	v.updateCount(s, cid)
}
