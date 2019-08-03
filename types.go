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
