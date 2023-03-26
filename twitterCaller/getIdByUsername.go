package twittercaller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const GET_ID_BY_USERNAME_URL = "https://api.twitter.com/2/users/by/username/:username"

type GetIdByUsernameResponse struct {
	Data struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
}

func (th *TwitterHandler) GetIdByUsername(username string) (string, error) {
	url := strings.Replace(GET_ID_BY_USERNAME_URL, ":username", username, 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on request construction.\n[ERROR] -", err)
		return "", err
	}
	req.Header.Add("Authorization", "bearer "+th.bearerToken)

	log.Println("GET request to " + url)
	// todo: set timeout on request
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return "", err
	}
	log.Println(string([]byte(body)))
	var responseObject GetIdByUsernameResponse
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", err
	}
	return responseObject.Data.Id, nil
}
