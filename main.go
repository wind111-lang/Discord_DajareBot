package main
/* Setup

   1. Create a .env file (click add file then remane it to .env)

   2. Put "token=" (without quotes) into the .env file followed by your Discord Bot token (No spaces!)

*/
import (
    "fmt"
    "os"
    // Uncomment below line if you are going to use uptimerobot to ping
    //"net/http"
    "github.com/bwmarrin/discordgo"

)
func main(){
    /* Uncomment this code block if you are going to use uptimerobot to ping
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    go http.ListenAndServe(":8080", nil)
    */
    // Create a new Discord session using the bot token from .env
    bot, err := discordgo.New("Bot " + os.Getenv("token"))

    if err != nil {
        panic(err)
    }

    // register events
    bot.AddHandler(ready)
    bot.AddHandler(messageCreate)

    err = bot.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
    // Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	for {}
	// Cleanly close down the Discord session.
	bot.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready){
    s.UpdateStatus(0, "with Go")
    fmt.Println("logged in as user " +string(s.State.User.ID))
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
    // Your code goes here

    if m.Content == "ping"{
        s.ChannelMessageSend(m.ChannelID, "pong")
    }
    
}