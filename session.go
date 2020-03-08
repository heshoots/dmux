package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordSession struct {
	*discordgo.Session
	handlers []Handler
}

func (s *DiscordSession) MessageChannel(c Channel, message Message) (Message, error) {
	m, err := s.Session.ChannelMessageSend(c.ID(), message.Content())
	return DiscordMessage{Message: m}, err
}

func (s *DiscordSession) AddHandler(h Handler) {
	s.handlers = append(s.handlers, h)
}

func (s *DiscordSession) Open() {
	s.Session.AddHandler(s.SessionRouter())
	s.Session.Open()
}

func (s *DiscordSession) RawSession() *discordgo.Session {
	return s.Session
}

/*
  Base router for all discord event types

  DiscordGo passes a variety of types to its message handlers

  When a message is recieved g will be a discordgo.MessageCreate object, the
  data from this can be added to the context that will be passed to upcoming
  handlers.

  The function then finds the handler that matches the event context and runs it
*/
func (s *DiscordSession) SessionRouter() func(*discordgo.Session, interface{}) {
	return func(d *discordgo.Session, g interface{}) {
		var handlerType HandlerType = -1
		context := DiscordGoContext{}
		switch g.(type) {
		case *discordgo.MessageCreate:
			handlerType = MessageHandler
			context.messageContext = DiscordMessageContext{
				g.(*discordgo.MessageCreate),
			}
		}
		for _, handler := range s.handlers {
			if handler.HandlerType() == handlerType {
				if handler.Pattern(context) {
					handler.Handle(&DiscordSession{d, []Handler{}}, context)
					return
				}
			}
		}
	}
}

func (s *DiscordSession) Close() {
	s.Session.Close()
}

func (s *DiscordSession) GuildRoles(g Guild) ([]Role, error) {
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

func (s *DiscordSession) UserPermissions(u User, c Channel) (Permissions, error) {
	p, err := s.Session.UserChannelPermissions(u.ID(), c.ID())
	return DiscordPermissions{p}, err
}

func (s *DiscordSession) GuildMemberRoleAdd(g Guild, u User, r Role) error {
	return s.Session.GuildMemberRoleAdd(g.ID(), u.ID(), r.ID())
}

func (s *DiscordSession) GuildRoleCreate(guild string, role string) (Role, error) {
	newrole, err := s.Session.GuildRoleCreate(guild)
	s.Session.GuildRoleEdit(guild, newrole.ID, role, 0, false, 0, false)
	return DiscordRole{Role: newrole}, err
}

func (s *DiscordSession) GuildMemberRoleRemove(g Guild, u User, r Role) error {
	return s.Session.GuildMemberRoleRemove(g.ID(), u.ID(), r.ID())
}
