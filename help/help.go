package help

import (
	"github.com/LotusJW/RLDBot/model"
	"github.com/bwmarrin/discordgo"
)

func Handle(s *discordgo.Session, m *discordgo.MessageCreate, parameters []string) {
	embed := discordgo.MessageEmbed{
		Title: "List of Commands",
		Color: 0xEF4923,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  model.Prefix + "chat",
				Value: "Send Rocket League quickchat messages. Requires two directions as a parameter, eg. ld (left down). Optionally, provide amount of times to send message, eg. 3. Example: " + model.Prefix + "chat ld 3",
			},
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

func UnknownMessage(s *discordgo.Session, m *discordgo.MessageCreate, parameters []string) {
	s.ChannelMessageSend(m.ChannelID, "I don't recognise that command, try "+model.Prefix+"help")
}
