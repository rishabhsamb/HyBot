package outbursts

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"

	"strconv"
)

type outburst struct {
	Key            string   `firestore:"key"`
	CallCount      int64    `firestore:"callCount"`
	Messages       []string `firestore:"messages"`
	RandomMessages []string `firestore:"randomMessages"`
}

func (v *outburst) getRandomMessageIndex() int {
	randomSeed := rand.NewSource(time.Now().UnixNano())
	randomEngine := rand.New(randomSeed)
	return randomEngine.Intn(len(v.RandomMessages))
}

func (v *outburst) sendMessages(s *discordgo.Session, cid string) {
	for _, str := range v.Messages {
		s.ChannelMessageSend(cid, str)
	}
	if len(v.RandomMessages) > 0 {
		s.ChannelMessageSend(cid, v.RandomMessages[v.getRandomMessageIndex()])
	}
}

func (v *outburst) updateCount(s *discordgo.Session, cid string) {
	v.CallCount++
	s.ChannelMessageSend(cid, v.Key+" has been called "+strconv.FormatInt(v.CallCount, 10)+" times.")
}

func (v *outburst) fire(s *discordgo.Session, cid string) {
	v.sendMessages(s, cid)
	v.updateCount(s, cid)
}
