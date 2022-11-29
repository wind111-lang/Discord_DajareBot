package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv"
	"github.com/kurehajime/dajarep"
)

//var sc = bufio.NewScanner(os.Stdin)

func sendMessage(s *discordgo.Session, channelID, message string) {
	err := s.MessageReactionAdd(channelID, message, "👀")
	if err != nil {
		fmt.Println(err)
	}
	_, err = s.ChannelMessageSend(channelID, message)
	log.Println(">>> " + message)
	if err != nil {
		log.Println("Error sending message or reaction: ", err)
	}
}

// func sendReply(s *discordgo.Session, channelID, message string, reference *discordgo.MessageReference) {
// 	_, err := s.ChannelMessageSendReply(channelID, message, reference)
// 	if err != nil {
// 		log.Fatal("Error sending message: ", err)
// 	}
// }

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	clientId := os.Getenv("CLIENT_ID")
	u := m.Author
	fmt.Printf("%20s %20s(%20s) > %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	if u.ID != clientId {
		dajare, _ := dajarep.Dajarep(m.Content)
		//fmt.Println(dajare)

		if len(dajare) > 0 {

			sendMessage(s, m.ChannelID, u.Mention()+"なんか言ってるねえ -> "+m.Content)
			//sendReply(s, m.ChannelID, "対象:"+m.Content, m.Reference())
		}
	}
}

func main() {

	//var a string
	//enverr := godotenv.Load(fmt.Sprint(".env"))
	// if enverr != nil {
	// 	panic("Error loading .env file")
	// }

	token := os.Getenv("DISCORD_TOKEN")

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}
	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening Discord: ", err)
	}

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot

	err = discord.Close()
	if err != nil {
		panic(err)
	}
	return
}
