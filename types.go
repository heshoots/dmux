package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type HandlerType int

const (
	MessageHandler = 0
)

type DiscordMessageContext struct {
	messageCreate *discordgo.MessageCreate
}

func (d DiscordMessageContext) GuildID() string {
	return d.messageCreate.GuildID
}

func (d DiscordMessageContext) ChannelID() string {
	return d.messageCreate.ChannelID
}

func (d DiscordMessageContext) Message() string {
	return d.messageCreate.Content
}

func (d DiscordMessageContext) UserID() string {
	return d.messageCreate.Author.ID
}

func (d DiscordMessageContext) UserName() string {
	return d.messageCreate.Author.Username
}

func (d DiscordMessageContext) UserAdmin(s Session) (bool, error) {
	permissions, err := s.UserChannelPermissions(d.UserID(), d.ChannelID())
	return permissions.isAdmin(), err
}

type MessageContext interface {
	GuildID() string
	ChannelID() string
	Message() string
	UserID() string
	UserName() string
	UserAdmin(s Session) (bool, error)
}

type HandlerContext interface {
	MessageContext() (bool, MessageContext)
}

type DiscordGoContext struct {
	messageContext MessageContext
}

func (d DiscordGoContext) MessageContext() (bool, MessageContext) {
	return d.messageContext != nil, d.messageContext
}

type Handler interface {
	Handle(c Session, context HandlerContext)
	HandlerType() HandlerType
	Name() string
	Description() string
	Pattern(HandlerContext) bool
}

type Message interface {
}

type Guild interface {
	Name() string
	ID() string
}

type Role interface {
	Name() string
	ID() string
}

type Permissions interface {
	isAdmin() bool
}

type DiscordPermissions struct {
	int
}

func (p DiscordPermissions) isAdmin() bool {
	var admin = discordgo.PermissionAdministrator
	return (p.int & admin) == admin
}

type Session interface {
	MessageChannel(channelID, message string) (Message, error)
	GuildRoles(guildID string) ([]Role, error)
	GuildMemberRoleAdd(guildID, userID, roleID string) error
	GuildMemberRoleRemove(guildID, userID, roleID string) error
	UserChannelPermissions(userID, channelID string) (Permissions, error)
	AddHandler(Handler)
	Open()
	Close()
}

type DiscordMessage struct {
	*discordgo.Message
}

type DiscordGuild struct {
	*discordgo.Guild
}

type DiscordRole struct {
	*discordgo.Role
}

type DiscordSession struct {
	*discordgo.Session
	handlers []Handler
}

func (s DiscordSession) MessageChannel(channelID string, message string) (Message, error) {
	m, err := s.Session.ChannelMessageSend(channelID, message)
	return DiscordMessage{m}, err
}

func (s DiscordSession) AddHandler(h Handler) {
	handlers = append(handlers, h)
}

var handlers []Handler

func DiscordRouter(s *discordgo.Session, g interface{}) {
	var handlerType HandlerType = -1
	context := DiscordGoContext{}
	switch g.(type) {
	case *discordgo.MessageCreate:
		handlerType = MessageHandler
		context.messageContext = DiscordMessageContext{
			g.(*discordgo.MessageCreate),
		}
	}
	for _, handler := range handlers {
		if handler.HandlerType() == handlerType {
			if handler.Pattern(context) {
				handler.Handle(DiscordSession{s, []Handler{}}, context)
				return
			}
		}
	}
}

func (s DiscordSession) Open() {
	s.Session.AddHandler(DiscordRouter)
	s.Session.Open()
}

func (s DiscordSession) Close() {
	s.Session.Close()
}

func (s DiscordSession) GuildRoles(guildID string) ([]Role, error) {
	r, err := s.Session.GuildRoles(guildID)
	if err != nil {
		return nil, err
	}
	roles := make([]Role, len(r))
	for i, role := range r {
		roles[i] = DiscordRole{role}
	}
	return roles, nil
}

func (s DiscordSession) UserChannelPermissions(userID, channelID string) (Permissions, error) {
	p, err := s.Session.State.UserChannelPermissions(userID, channelID)
	return DiscordPermissions{p}, err
}

func (s DiscordRole) ID() string {
	return s.Role.ID
}

func (s DiscordRole) Name() string {
	return s.Role.Name
}
