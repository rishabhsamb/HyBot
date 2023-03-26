package twittercaller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	TWEET_LOOKUP_URL = "https://api.twitter.com/2/users/:id/tweets?max_results=:num&since_id=:since_id&exclude=retweets,replies"
)

type Tweet struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type GetMostRecentTweetsResponse struct {
	Data []Tweet `json:"data"`
}

// THIS METHOD MUST BE EXECUTED UNDERNEATH A LOCK
func (th *TwitterHandler) GetMostRecentTweets(username string, num int64, sinceId string) ([]Tweet, error) {
	if num > 25 || num < 5 {
		return nil, errors.New("num out of acceptable range, must be an integer in [5, 25]")
	}

	var userId string
	if _, ok := th.tweetSubscriptions[username]; ok {
		userId = th.tweetSubscriptions[username].UserId
	} else {
		getUserId, err := th.GetIdByUsername(username)
		if err != nil {
			log.Println("Error on getting ID by username.\n[ERROR] -", err)
			return nil, err
		}
		userId = getUserId
	}

	url := strings.Replace(strings.Replace(strings.Replace(TWEET_LOOKUP_URL, ":id", userId, 1), ":num", strconv.FormatInt(num, 10), 1), ":since_id", sinceId, 1)
	log.Println("url is " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on request construction.\n[ERROR] -", err)
		return nil, err
	}
	req.Header.Add("Authorization", "bearer "+th.bearerToken)

	log.Println("GET request to " + url)
	// todo: set timeout on request
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes.\n[ERROR] -", err)
		return nil, err
	}
	log.Println("body of response: ")
	log.Println(string([]byte(body)))
	var responseObject GetMostRecentTweetsResponse
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return nil, err
	}
	return responseObject.Data, nil
}
