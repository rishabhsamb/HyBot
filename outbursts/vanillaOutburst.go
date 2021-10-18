package outbursts

import (
	"github.com/bwmarrin/discordgo"

	"strconv"
)

type vanillaOutburst struct {
	key       string
	callCount uint64
	messages  []string
}

func (v *vanillaOutburst) getKey() string {
	return v.key
}

func (v *vanillaOutburst) sendMessages(s *discordgo.Session, cid string) {
	for _, str := range v.messages {
		s.ChannelMessageSend(cid, str)
	}
}

func (v *vanillaOutburst) updateCount(s *discordgo.Session, cid string) {
	v.callCount++
	s.ChannelMessageSend(cid, v.key+" has been called "+strconv.FormatUint(v.callCount, 10)+" times.")
}

func (v *vanillaOutburst) fire(s *discordgo.Session, cid string) {
	v.sendMessages(s, cid)
	v.updateCount(s, cid)
}

func (v *vanillaOutburst) load() {

}

func (v *vanillaOutburst) add() {

}
