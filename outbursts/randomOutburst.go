package outbursts

import (
	"math/rand"

	"time"

	"github.com/bwmarrin/discordgo"
)

type randomOutburst struct {
	base           vanillaOutburst
	randomMessages []string
}

func (ro *randomOutburst) getRandomMessageIndex() int {
	randomSeed := rand.NewSource(time.Now().UnixNano())
	randomEngine := rand.New(randomSeed)
	return randomEngine.Intn(len(ro.randomMessages))
}

func (r *randomOutburst) fire(s *discordgo.Session, cid string) {
	r.base.sendMessages(s, cid)
	randomIndex := r.getRandomMessageIndex()
	s.ChannelMessageSend(cid, r.randomMessages[randomIndex])
	r.base.updateCount(s, cid)
}

func (r *randomOutburst) load() {

}

func (r *randomOutburst) add() {

}
