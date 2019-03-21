package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordChannel struct {
	id      string
	channel *discordgo.Channel
}

func (c DiscordChannel) ID() string {
	return c.id
}
