package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordSession struct {
	*discordgo.Session
	handlers []Handler
}

var handlers []Handler

func (s DiscordSession) MessageChannel(c Channel, message Message) (Message, error) {
	m, err := s.Session.ChannelMessageSend(c.ID(), message.Content())
	return DiscordMessage{Message: m}, err
}

func (s DiscordSession) AddHandler(h Handler) {
	handlers = append(handlers, h)
}

func (s DiscordSession) Open() {
	s.Session.AddHandler(DiscordRouter)
	s.Session.Open()
}

func (s DiscordSession) Close() {
	s.Session.Close()
}

func (s DiscordSession) GuildRoles(g Guild) ([]Role, error) {
	r, err := s.Session.GuildRoles(g.ID())
	if err != nil {
		return nil, err
	}
	roles := make([]Role, len(r))
	for i, role := range r {
		roles[i] = DiscordRole{Role: role}
	}
	return roles, nil
}

func (s DiscordSession) UserPermissions(u User, c Channel) (Permissions, error) {
	p, err := s.Session.UserChannelPermissions(u.ID(), c.ID())
	return DiscordPermissions{p}, err
}

func (s DiscordSession) GuildMemberRoleAdd(g Guild, u User, r Role) error {
	return s.Session.GuildMemberRoleAdd(g.ID(), u.ID(), r.ID())
}

func (s DiscordSession) GuildMemberRoleRemove(g Guild, u User, r Role) error {
	return s.Session.GuildMemberRoleAdd(g.ID(), u.ID(), r.ID())
}
