package dmux

import (
	"github.com/bwmarrin/discordgo"
)

// Provides information about an incoming message to a handler
type MessageContext interface {
	Guild() Guild
	Channel() Channel
	Message() Message
	User() User
}

/*
  Provides relevant contextual information to a handler

  a handler which responds to messages would use MessageContext to obtain
  the message information
*/
type HandlerContext interface {
	MessageContext() (bool, MessageContext)
}

/*
  A handler which responds to certain router commands

  Currently supported handler types:
	MessageHandler
*/
type Handler interface {
	Handle(c Session, context HandlerContext)
	HandlerType() HandlerType
	Name() string
	Description() string
	Pattern(HandlerContext) bool
}

type Message interface {
	ID() string
	Content() string
	Author() User
}

type User interface {
	ID() string
	Name() string
	Admin(s Session, c Channel) (bool, error)
}

type Channel interface {
	ID() string
}

type Guild interface {
	ID() string
	Roles(s Session) ([]Role, error)
}

type Role interface {
	Name() string
	ID() string
}

type Permissions interface {
	isAdmin() bool
}

type Session interface {
	// Send a message to a channel
	MessageChannel(c Channel, m Message) (Message, error)
	// Get list of roles for a joined guild
	GuildRoles(g Guild) ([]Role, error)
	// Add role to a user within a guild
	GuildMemberRoleAdd(g Guild, u User, r Role) error
	// Remove a role from a user within a guild
	GuildMemberRoleRemove(g Guild, u User, r Role) error
	// Create a role with given role name in guild
	GuildRoleCreate(guild string, role string) (Role, error)
	// Get permissions for a user in a channel
	UserPermissions(u User, c Channel) (Permissions, error)
	// Add handler to router
	AddHandler(Handler)
	Open()
	Close()
	RawSession() *discordgo.Session
}
