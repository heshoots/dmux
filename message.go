package dmux

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordMessage struct {
	*discordgo.Message
	content string
}

func DiscordMessageString(content string) DiscordMessage {
	return DiscordMessage{content: content}
}

func (m DiscordMessage) Content() string {
	if m.Message != nil {
		return m.Message.Content
	} else {
		return m.content
	}
}

func (m DiscordMessage) ID() string {
	if m.Message != nil {
		return m.Message.ID
	} else {
		panic("Message does not have an ID")
	}
}

func (m DiscordMessage) Author() User {
	if m.Message != nil {
		return DiscordUser{m.Message.Author}
	} else {
		panic("Message does not have user")
	}
}
