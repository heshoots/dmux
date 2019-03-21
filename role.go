package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordRole struct {
	*discordgo.Role
	id string
}

func CreateDiscordRole(id string) Role {
	return DiscordRole{Role: nil, id: id}
}

func (s DiscordRole) ID() string {
	return s.Role.ID
}

func (s DiscordRole) Name() string {
	return s.Role.Name
}
