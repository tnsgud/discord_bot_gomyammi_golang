package cmd

import "github.com/bwmarrin/discordgo"

func GoogleImageController(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")

	}
}
