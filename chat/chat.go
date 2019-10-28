package chat

import (
  "time"
  "strconv"

  "github.com/bwmarrin/discordgo"
  "github.com/LotusJW/RLDBot/config"
  "github.com/LotusJW/RLDBot/model"
)

const (
  defaultAmount = 5
  cooldown = 4
  delay = time.Millisecond * 200
)

func Handle(s *discordgo.Session, m *discordgo.MessageCreate, parameters []string) {
  if len(parameters) < 2 {
    noParams(s, m)
    return
  }

  messageID := parameters[1]
  if len(messageID) != 2 {
    notFound(s, m, messageID)
    return
  }

  firstID := string(messageID[0])
  secondID := string(messageID[1])

  if _, ok := config.Messages[firstID]; !ok {
    notFound(s, m, messageID)
    return
  }

  if _, ok := config.Messages[firstID][secondID]; !ok {
    notFound(s, m, messageID)
    return
  }

  message := config.Messages[firstID][secondID]

  amount := defaultAmount
  if len(parameters) >= 3 {
    i, err := strconv.Atoi(parameters[2])
    if err != nil {
      invalidAmount(s, m)
      return
    }

    amount = i
  }

  sendMessages(s, m, message, amount)
}

func sendMessages(s *discordgo.Session, m *discordgo.MessageCreate, message string, amount int) {
  for i := 0; i < amount; i++ {
    if i == cooldown {
      s.ChannelMessageSend(m.ChannelID, "Chat disabled for 3 seconds")
      break
    }

    s.ChannelMessageSend(m.ChannelID, message)
    time.Sleep(delay)
  }
}

func notFound(s *discordgo.Session, m *discordgo.MessageCreate, messageID string) {
  s.ChannelMessageSend(m.ChannelID, "I don't recognise the message " + messageID + ", try " + model.Prefix + "help")
}

func noParams(s *discordgo.Session, m *discordgo.MessageCreate) {
  s.ChannelMessageSend(m.ChannelID, "No parameters provided, try " + model.Prefix + "help")
}

func invalidAmount(s *discordgo.Session, m *discordgo.MessageCreate) {
  s.ChannelMessageSend(m.ChannelID, "The amount provided is invalid, try " + model.Prefix + "help")
}
