package twittercaller

import (
	"errors"
	"sync"
)

type SubscriptionKind int

const (
	Like  SubscriptionKind = iota
	Tweet SubscriptionKind = iota
)

type User struct {
	UserId        string
	Subscriptions map[SubscriptionKind]bool
}

type TwitterHandler struct {
	mu    sync.RWMutex
	users map[string]User
}

// func (th *TwitterHandler) Subscribe(username string, kind SubscriptionKind) error {
// 	th.mu.Lock()
// 	defer th.mu.Unlock()

// 	if _, ok := th.users[username]; ok {
// 		th.users[username].Subscriptions[kind] = true
// 		return nil
// 	} else {
// 		userId, err := th.getIdByUsername(username)
// 		if err {
// 			return err
// 		}
// 		subscriptions := make(map[SubscriptionKind]bool)
// 		subscriptions[kind] = true
// 		th.users[username] = User{userId, subscriptions}
// 		return nil
// 	}
// }

func (th *TwitterHandler) Unsubscribe(username string, kind SubscriptionKind) error {
	th.mu.Lock()
	defer th.mu.Unlock()

	if _, ok := th.users[username]; ok {
		th.users[username].Subscriptions[kind] = false
		return nil
	} else {
		return errors.New("there is no current subscription to " + username)
	}
}

// func (th *TwitterHandler) GetMostRecentTweet(username string) error

// func (th *TwitterHandler) GetMostRecentLike(username string) error
