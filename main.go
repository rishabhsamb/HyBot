package main

// import (
// 	"fmt"
// 	"log"

// 	twittercaller "github.com/rishabhsamb/HyBot/twitterCaller"
// )

// func main() {
// 	th := new(twittercaller.TwitterHandler)
// 	err := th.Register()
// 	if err != nil {
// 		log.Println("Error on twitter handler registration.\n[ERROR] -", err)
// 		return
// 	}
// 	// log.Println(strings.Replace("https://api.twitter.com/2/users/by/username/:username", ":username", "cock", 1))
// 	tweets, err := th.GetMostRecent("zephanijong", 10, twittercaller.TweetKind)
// 	if err != nil {
// 		log.Println("Error on sourcing tweets.\n[ERROR] -", err)
// 	}
// 	for _, val := range tweets {
// 		fmt.Println(val.Text)
// 	}
// }

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	twittercaller "github.com/rishabhsamb/HyBot/twitterCaller"
)

const (
	DISCUSSION_CHANNEL_ID = "754877121889042443"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	// ctx := context.Background()
	// firestoreClient := createClient(ctx)
	var (
		// commander commandHandler
		th twittercaller.TwitterHandler = twittercaller.TwitterHandler{}
	)

	// commander.init(ctx, firestoreClient)

	token := os.Getenv("DISCORD_TOKEN")
	b, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err.Error())
	}
	// b.AddHandler(commander.driver)

	err = b.Open()
	if err != nil {
		log.Panic("Could not connect to Discord", err)
		return
	}
	subQuitChan := make(chan bool)
	err = th.Register(b, subQuitChan, DISCUSSION_CHANNEL_ID)
	if err != nil {
		log.Panic("Error on registering Twitter handler.\n[ERROR] -", err)
		return
	}

	defer b.Close()

	log.Print("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
