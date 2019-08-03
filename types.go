package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type HandlerType int

const (
	// A handler which consumers messages
	MessageHandler = 0
)

type DiscordGoContext struct {
	messageContext MessageContext
}

func (d DiscordGoContext) MessageContext() (bool, MessageContext) {
	return d.messageContext != nil, d.messageContext
}

type DiscordPermissions struct {
	int
}

func (p DiscordPermissions) isAdmin() bool {
	var admin = discordgo.PermissionAdministrator
	return (p.int & admin) == admin
}
