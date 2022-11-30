package main

import (
	"fmt"
	"log"
	_ "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv"
	"github.com/kurehajime/dajarep"
)

//var sc = bufio.NewScanner(os.Stdin)

func sendMessage(s *discordgo.Session, channelID, message string) {
	//var r *discordgo.MessageReactionAdd
	_, err := s.ChannelMessageSend(channelID, message)
	log.Println(">>> " + message)
	if err != nil {
		log.Println("Error sending message or : ", err)
	}
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	u := m.Author
	fmt.Printf("%20s %20s(%20s) > %s\n", m.ChannelID, u.Username, u.ID, m.Content)
	clientId := os.Getenv("CLIENT_ID")
	if u.ID != clientId {
		dajare, _ := dajarep.Dajarep(m.Content)
		//fmt.Println(dajare)

		if len(dajare) > 0 {
			err := s.MessageReactionAdd(m.ChannelID, m.Message.ID, "👀")
			if err != nil {
				log.Println(err)
			}
			sendMessage(s, m.ChannelID, u.Mention()+"なんか言ってるねえ -> "+m.Content)
		}
	}
}

func main() {
	token := os.Getenv("DISCORD_TOKEN")

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}
	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening Discord session!: ", err)
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

// func main() {
// 	http.HandleFunc("/", bot)
// 	http.ListenAndServe(":8080", nil)
// }
