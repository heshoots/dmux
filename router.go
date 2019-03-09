package dmux

import (
	"github.com/bwmarrin/discordgo"
)

func Router(authToken string) (Session, error) {
	d, err := discordgo.New("Bot " + authToken)
	if err != nil {
		return nil, err
	}
	return DiscordSession{d, []Handler{}}, nil
}
