package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordUser struct {
	*discordgo.User
}

func (u DiscordUser) Name() string {
	return u.User.Username
}

func (u DiscordUser) ID() string {
	return u.User.ID
}

func (u DiscordUser) Admin(s Session, c Channel) (bool, error) {
	p, err := s.UserPermissions(u, c)
	if err != nil {
		return false, err
	}
	return p.isAdmin(), nil
}
