package discord

import (
	"github.com/bwmarrin/discordgo"
)

func HyoonHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hyoon" {
		s.ChannelMessageSend(m.ChannelID, "hyoon")
		return
	}
}
