package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordGuild struct {
	id string
	*discordgo.Guild
}

func (g DiscordGuild) ID() string {
	return g.id
}

func (g DiscordGuild) Roles(s Session) ([]Role, error) {
	return s.GuildRoles(g)
}
