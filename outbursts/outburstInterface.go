package outbursts

import (
	"github.com/bwmarrin/discordgo"
)

type outburst interface {
	fire(s *discordgo.Session, cid string)
	getKey() string
	save() error
}
