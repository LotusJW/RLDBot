package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/LotusJW/RLDBot/chat"
	"github.com/LotusJW/RLDBot/help"
	"github.com/LotusJW/RLDBot/model"
	"github.com/LotusJW/RLDBot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	token string
)

// go run main.go -t THE_DISCORD_TOKEN
func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	if token == "" {
		log.Println("Error: no token provided")
		return
	}

	err := config.Load()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Printf("Error creating Discord session: %v", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Printf("Error opening connection: %v", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.UpdateStatus(0, model.Prefix+"help")
	if err != nil {
		log.Printf("Error updating status: %v", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("RLDBot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, model.Prefix) {
		return
	}

	// Original message: "?chat ld 5".

	// Remove the prefix.
	// Command is now "chat ld 5".
	command := strings.TrimPrefix(m.Content, model.Prefix)

	// Split the message between spaces into an array.
	// Parameters are now "chat", "ld", "5"
	parameters := strings.Split(command, " ")

	// Get the first parameter which is the command.
	// First parameter is "chat".
	command = parameters[0]

	switch command {
	case "chat":
		chat.Handle(s, m, parameters)

	case "help":
		help.Handle(s, m, parameters)

	default:
		help.UnknownMessage(s, m, parameters)
	}
}
