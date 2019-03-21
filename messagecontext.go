package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordMessageContext struct {
	messageCreate *discordgo.MessageCreate
}

func (d DiscordMessageContext) Guild() Guild {
	return DiscordGuild{id: d.messageCreate.GuildID}
}

func (d DiscordMessageContext) Channel() Channel {
	return DiscordChannel{id: d.messageCreate.ChannelID}
}

func (d DiscordMessageContext) Message() Message {
	return DiscordMessage{Message: d.messageCreate.Message}
}

func (d DiscordMessageContext) User() User {
	return DiscordUser{d.messageCreate.Author}
}

func (d DiscordMessageContext) UserName() string {
	return d.messageCreate.Author.Username
}

func (d DiscordMessageContext) UserAdmin(s Session) (bool, error) {
	permissions, err := s.UserPermissions(d.User(), d.Channel())
	return permissions.isAdmin(), err
}
