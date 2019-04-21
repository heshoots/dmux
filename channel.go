package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordChannel struct {
	*discordgo.Channel
	id string
}

func CreateDiscordChannel(id string) DiscordChannel {
	return DiscordChannel{Channel: nil, id: id}
}

func (c DiscordChannel) ID() string {
	if c.Channel != nil {
		return c.Channel.ID
	} else {
		return c.id
	}
}
