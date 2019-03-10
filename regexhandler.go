package dmux

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"regexp"
)

type RegexMessageHandler interface {
	Handle(c *discordgo.Session, context HandlerContext)
	Pattern(context HandlerContext) bool
	HandlerType() HandlerType
	Name() string
	Description() string
	NeedsAdmin() bool
}

type DiscordRegexMessageHandler struct {
	HandlerPattern     string
	HandlerFn          func(c Session, context RegexHandlerContext)
	HandlerName        string
	HandlerDescription string
	RequiresAdmin      bool
}

type RegexHandlerContext interface {
	MessageContext() (bool, MessageContext)
	Groups() map[string]string
}

type DiscordRegexHandlerContext struct {
	HandlerContext
	pattern string
}

func (h DiscordRegexMessageHandler) NeedsAdmin() bool {
	return h.RequiresAdmin
}

func (h DiscordRegexMessageHandler) Description() string {
	return h.HandlerDescription
}

func (h DiscordRegexHandlerContext) Groups() map[string]string {
	ok, ctx := h.MessageContext()
	matchMap := make(map[string]string)
	if ok {
		regex := regexp.MustCompile(h.pattern)
		groups := regex.FindStringSubmatch(ctx.Message())
		names := regex.SubexpNames()
		for i, name := range names {
			matchMap[name] = groups[i]
		}
	}
	return matchMap
}

func (h *DiscordRegexMessageHandler) Handle(s Session, context HandlerContext) {
	if h.Pattern(context) {
		ok, ctx := context.MessageContext()
		if !ok {
			log.Error("Couldn't read message context")
			return
		}
		isAdmin, err := ctx.UserAdmin(s)
		if err != nil {
			log.Error("Couldn't determine if admin user")
			return
		}
		if h.NeedsAdmin() && !isAdmin {
			log.Error("User not admin")
			s.MessageChannel(ctx.ChannelID(), "Only admins can do this")
			return
		}
		log.WithFields(log.Fields{
			"Handler":     h.Name(),
			"Description": h.Description(),
			"User_ID":     ctx.UserID(),
			"User_Name":   ctx.UserName(),
		}).Info()
		h.HandlerFn(s, DiscordRegexHandlerContext{
			context,
			h.HandlerPattern,
		})
	}
}

func (h *DiscordRegexMessageHandler) Pattern(context HandlerContext) bool {
	ok, ctx := context.MessageContext()
	if ok {
		regex := regexp.MustCompile(h.HandlerPattern)
		return regex.MatchString(ctx.Message())
	}
	return false
}

func (h *DiscordRegexMessageHandler) HandlerType() HandlerType {
	return MessageHandler
}

func (h *DiscordRegexMessageHandler) Name() string {
	return h.HandlerName
}
