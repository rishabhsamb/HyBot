package twittercaller

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type TweetRelationshipKind int

const (
	LikeKind  TweetRelationshipKind = iota
	TweetKind TweetRelationshipKind = iota
)

type User struct {
	Username string
	UserId   string
	SinceId  string
}

type TwitterHandler struct {
	mu                 sync.RWMutex
	tweetSubscriptions map[string]User
	bearerToken        string
}

func (th *TwitterHandler) Register(s *discordgo.Session, quitChan chan bool, subscriptionSendChannel string) error {
	th.mu = sync.RWMutex{}
	th.mu.Lock()
	defer th.mu.Unlock()
	th.tweetSubscriptions = make(map[string]User)
	bearerToken, ok := os.LookupEnv("TWITTER_BEARER_TOKEN")
	if ok {
		th.bearerToken = bearerToken
		go th.sendTweetSubscriptions(s, quitChan, subscriptionSendChannel)
		return nil
	} else {
		return errors.New("could not find TWITTER_BEARER_TOKEN in environment variables")
	}
}

func (th *TwitterHandler) Subscribe(username string) error {
	th.mu.RLock()

	if _, ok := th.tweetSubscriptions[username]; ok {
		th.mu.RUnlock()
		return nil
	} else {
		th.mu.RUnlock()
		userId, err := th.GetIdByUsername(username)
		if err != nil {
			log.Println("Error on getting ID by username.\n[ERROR] -", err)
			return err
		}
		th.mu.Lock()
		th.tweetSubscriptions[username] = User{username, userId, "20"}
		th.mu.Unlock()
		return nil
	}
}

func (th *TwitterHandler) Unsubscribe(username string) error {
	th.mu.Lock()
	defer th.mu.Unlock()

	if _, ok := th.tweetSubscriptions[username]; ok {
		delete(th.tweetSubscriptions, username)
		return nil
	} else {
		return errors.New("there is no current subscription to " + username)
	}
}

func (th *TwitterHandler) sendTweetSubscriptions(s *discordgo.Session, quitChan chan bool, subscriptionSendChannel string) {
	s.ChannelMessageSend(subscriptionSendChannel, "update routine starting")
	for {
		select {
		case <-quitChan:
			s.ChannelMessageSend(subscriptionSendChannel, "quitting subscription updates")
			return
		case <-time.After(time.Second * 10):
			th.mu.Lock()
			log.Println("updating tweets")
			for username, user := range th.tweetSubscriptions {
				tweets, err := th.GetMostRecentTweets(username, 5, user.SinceId)
				if err != nil {
					log.Println("could not update " + username + "'s " + "tweets")
				} else if len(tweets) != 0 {
					for _, tweet := range tweets {
						s.ChannelMessageSend(subscriptionSendChannel, getTweetLink(tweet.Id, username))
					}
					user.SinceId = tweets[0].Id
					th.tweetSubscriptions[username] = user
				}
			}
			th.mu.Unlock()
		}
	}
}

func getTweetLink(tweetId string, tweetUsername string) string {
	return "https://twitter.com/" + tweetUsername + "/status/" + tweetId
}
